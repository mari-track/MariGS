package Game

import (
	"github.com/mari-track/MariGS/internal/Game/model"
	"github.com/mari-track/MariGS/pkg/gdconf"
	"github.com/mari-track/MariGS/protocol/cmd"
	"github.com/mari-track/MariGS/protocol/proto"
)

func (g *Game) AddAvatar(avatarId uint32) {
	conf := gdconf.GetAvatarDataById(avatarId)
	if conf == nil {
		return
	}
	weaponGuid := uint64(model.SNOWFLAKE.GenId())
	// 先添加角色
	g.Player.AddAvatar(avatarId)
	// 然后添加武器
	g.AddWeapon(weaponGuid, conf.InitialWeapon)

	// 发送通知
	g.AvatarFightPropUpdateNotify(avatarId)
}

func (g *Game) AddWeapon(guid uint64, weaponId uint32) {
	g.Player.AddWeapon(guid, weaponId)
	// 添加通知
	guidList := []uint64{guid}
	g.WeaponStoreItemChangeNotify(guidList)
}

// 角色装备武器
func (g *Game) UpdateAvatarWeapon(avatarId uint32, guid uint64) {
	// 修改角色属性武器
	g.Player.UpdateAvatarWeapon(guid, avatarId)
	// 修改武器属性角色
	g.Player.UpdateWeaponAvatar(guid, avatarId)
}

/*************************************通知包******************************************/

// 武器通知
func (g *Game) WeaponStoreItemChangeNotify(guidList []uint64) {
	notify := &proto.StoreItemChangeNotify{
		StoreType: proto.StoreType_STORE_PACK,
	}
	for _, guid := range guidList {
		db := g.Player.GetPbWeaponByGuid(guid)
		if db == nil {
			continue
		}
		item := &proto.Item{
			ItemId: db.WeaponId,
			Guid:   db.Guid,
			Detail: &proto.Item_Equip{
				Equip: &proto.Equip{
					Detail: &proto.Equip_Weapon{
						Weapon: &proto.Weapon{
							Level:        db.Level,
							Exp:          db.Exp,
							PromoteLevel: db.PromoteLevel,
							AffixMap:     db.AffixMap,
						},
					},
				},
			},
		}
		notify.ItemList = append(notify.ItemList, item)
	}
	g.seed(cmd.StoreItemChangeNotify, notify)
}
