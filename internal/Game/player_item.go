package Game

import (
	"github.com/mari-track/MariGS/internal/Game/model"
	"github.com/mari-track/MariGS/pkg/constant"
	"github.com/mari-track/MariGS/pkg/gdconf"
	"github.com/mari-track/MariGS/protocol/cmd"
	"github.com/mari-track/MariGS/protocol/proto"
)

/*************************************通知包******************************************/

func (g *Game) StoreWeightLimitNotify() {
	notify := &proto.StoreWeightLimitNotify{
		StoreType: proto.StoreType_STORE_PACK,
		// 背包容量限制
		WeightLimit: constant.STORE_PACK_LIMIT_WEIGHT,
	}
	g.seed(cmd.StoreWeightLimitNotify, notify)
}

func (g *Game) PlayerStoreNotify() {
	notify := &proto.PlayerStoreNotify{
		StoreType:   proto.StoreType_STORE_PACK,
		WeightLimit: constant.STORE_PACK_LIMIT_WEIGHT,
		ItemList:    make([]*proto.Item, 0),
	}
	// add 武器
	weaponList := g.Player.GetPbWeaponList()
	for _, weapon := range weaponList {
		item := &proto.Item{
			ItemId: weapon.WeaponId,
			Guid:   weapon.Guid,
			Detail: &proto.Item_Equip{
				Equip: &proto.Equip{
					Detail: &proto.Equip_Weapon{
						Weapon: &proto.Weapon{
							Level:        weapon.Level,
							Exp:          weapon.Exp,
							PromoteLevel: weapon.PromoteLevel,
							AffixMap:     weapon.AffixMap,
						},
					},
				},
			},
		}
		notify.ItemList = append(notify.ItemList, item)
	}

	g.seed(cmd.PlayerStoreNotify, notify)
}

func (g *Game) ItemCdGroupTimeNotify() {
	notify := &proto.ItemCdGroupTimeNotify{
		ItemCdMap: make(map[uint32]uint64),
	}
	g.seed(cmd.ItemCdGroupTimeNotify, notify)
}

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

/*************************************封装******************************************/

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
