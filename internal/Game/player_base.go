package Game

import (
	"reflect"

	"github.com/mari-track/MariGS/pkg/logger"
	"github.com/mari-track/MariGS/protocol/cmd"
	"github.com/mari-track/MariGS/protocol/proto"
)

/*************************************通知包******************************************/

func (g *Game) PlayerDataNotify() {
	db := g.Player.GetPbPlayerBasicCompBin()
	notify := &proto.PlayerDataNotify{
		NickName:          db.Nickname,
		ServerTime:        uint64(GetServerTime()),
		IsFirstLoginToday: true,
		RegionId:          1,
		PropMap:           make(map[uint32]*proto.PropValue),
	}

	for k, v := range db.PropMap {
		notify.PropMap[k] = g.PacketPropValue(k, v)
	}
	g.seed(cmd.PlayerDataNotify, notify)
}

/*************************************封装******************************************/

func (g *Game) PacketPropValue(key uint32, value any) *proto.PropValue {
	propValue := new(proto.PropValue)
	propValue.Type = key
	switch value.(type) {
	case int:
		v := value.(int)
		propValue.Val = int64(v)
		propValue.Value = &proto.PropValue_Ival{Ival: int64(v)}
	case int8:
		v := value.(int8)
		propValue.Val = int64(v)
		propValue.Value = &proto.PropValue_Ival{Ival: int64(v)}
	case int16:
		v := value.(int16)
		propValue.Val = int64(v)
		propValue.Value = &proto.PropValue_Ival{Ival: int64(v)}
	case int32:
		v := value.(int32)
		propValue.Val = int64(v)
		propValue.Value = &proto.PropValue_Ival{Ival: int64(v)}
	case int64:
		v := value.(int64)
		propValue.Val = v
		propValue.Value = &proto.PropValue_Ival{Ival: v}
	case float32:
		v := value.(float32)
		propValue.Value = &proto.PropValue_Fval{Fval: v}
	case float64:
		v := value.(float64)
		propValue.Value = &proto.PropValue_Fval{Fval: float32(v)}
	case uint8:
		v := value.(uint8)
		propValue.Val = int64(v)
		propValue.Value = &proto.PropValue_Ival{Ival: int64(v)}
	case uint16:
		v := value.(uint16)
		propValue.Val = int64(v)
		propValue.Value = &proto.PropValue_Ival{Ival: int64(v)}
	case uint32:
		v := value.(uint32)
		propValue.Val = int64(v)
		propValue.Value = &proto.PropValue_Ival{Ival: int64(v)}
	case uint64:
		v := value.(uint64)
		propValue.Val = int64(v)
		propValue.Value = &proto.PropValue_Ival{Ival: int64(v)}
	default:
		logger.Error("unknown value type: %v, value: %v", reflect.TypeOf(value), value)
		return nil
	}
	return propValue
}
