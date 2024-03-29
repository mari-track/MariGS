package alg

import (
	"encoding/binary"
	"log"

	"github.com/mari-track/MariGS/pkg/endec"
)

// HK4E游戏协议编解码

/*
							《原神》KCP协议(带*为xor加密数据)
0			1			2					4											16(字节)
+---------------------------------------------------------------------------------------+
|					sessionId(le)			|					conv(le)				|
+---------------------------------------------------------------------------------------+
|	cmd		|	frg		|		wnd			|					ts						|
+---------------------------------------------------------------------------------------+
|						sn					|					una						|
+---------------------------------------------------------------------------------------+
|						len					|   	0x4567*		|		cmdId*			|
+---------------------------------------------------------------------------------------+
|		headLen*		|				payloadLen*				|		head*			|
+---------------------------------------------------------------------------------------+
|								payload*						|		 0x89AB*		|
+---------------------------------------------------------------------------------------+
*/

const (
	PacketMaxLen = 343 * 1024 // 最大应用层包长度
)

type PackMsg struct {
	CmdId     uint16
	HeadData  []byte
	ProtoData []byte
}

func DecodeBinToPayload(data []byte, kcpMsgList *[]*PackMsg, xorKey []byte) {
	// xor解密
	if xorKey != nil {
		endec.Xor(data, xorKey)
	}
	DecodeLoop(data, kcpMsgList)
	return
}

func DecodeLoop(data []byte, kcpMsgList *[]*PackMsg) {
	// 长度太短
	if len(data) < 12 {
		log.Println("packet len less than 12 byte")
		return
	}
	// 头部幻数错误
	if data[0] != 0x45 || data[1] != 0x67 {
		log.Println("packet head magic 0x4567 error")
		return
	}

	// 协议号
	cmdId := binary.BigEndian.Uint16(data[2:4])
	// 头部长度
	headLen := binary.BigEndian.Uint16(data[4:6])
	// proto长度
	protoLen := binary.BigEndian.Uint32(data[6:10])
	// 检查长度
	packetLen := int(headLen) + int(protoLen) + 12
	if packetLen > PacketMaxLen {
		log.Println("packet len too long")
		return
	}
	if len(data) < packetLen {
		log.Println("packet len not enough")
		return
	}
	// 尾部幻数错误
	if data[int(headLen)+int(protoLen)+10] != 0x89 || data[int(headLen)+int(protoLen)+11] != 0xAB {
		log.Println("packet tail magic 0x89AB error")
		return
	}
	// 头部数据
	headData := data[10 : 10+int(headLen)]
	// proto数据
	protoData := data[10+int(headLen) : 10+int(headLen)+int(protoLen)]
	// 返回数据
	kcpMsg := new(PackMsg)
	kcpMsg.CmdId = cmdId
	kcpMsg.HeadData = headData
	kcpMsg.ProtoData = protoData
	*kcpMsgList = append(*kcpMsgList, kcpMsg)
	// 有不止一个包 递归解析
	if len(data) > packetLen {
		DecodeLoop(data[packetLen:], kcpMsgList)
	}
}

func EncodePayloadToBin(kcpMsg *PackMsg, xorKey []byte) (bin []byte) {
	if kcpMsg.HeadData == nil {
		kcpMsg.HeadData = make([]byte, 0)
	}
	if kcpMsg.ProtoData == nil {
		kcpMsg.ProtoData = make([]byte, 0)
	}
	// 检查长度
	packetLen := len(kcpMsg.HeadData) + len(kcpMsg.ProtoData) + 12
	if packetLen > PacketMaxLen {
		log.Println("packet len too long")
		return make([]byte, 0)
	}
	bin = make([]byte, packetLen)
	// 头部幻数
	bin[0] = 0x45
	bin[1] = 0x67

	// 协议号
	binary.BigEndian.PutUint16(bin[2:4], kcpMsg.CmdId)
	// 头部长度
	binary.BigEndian.PutUint16(bin[4:6], uint16(len(kcpMsg.HeadData)))
	// proto长度
	binary.BigEndian.PutUint32(bin[6:10], uint32(len(kcpMsg.ProtoData)))
	// 头部数据
	copy(bin[10:], kcpMsg.HeadData)
	// proto数据
	copy(bin[10+len(kcpMsg.HeadData):], kcpMsg.ProtoData)
	// 尾部幻数
	bin[len(bin)-2] = 0x89
	bin[len(bin)-1] = 0xAB
	// xor加密
	if xorKey != nil {
		endec.Xor(bin, xorKey)
	}
	return bin
}
