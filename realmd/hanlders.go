package realmd

import (
	"mircore/game"
	"mircore/game/proto"
	"mircore/utils/log"
	"strings"
)

//handleIncoming process incoming packet
func handleIncoming(s game.Session, packet *proto.WorldPacket) int {
	ret := 0

	switch proto.OpcodeType(packet.Opcode) {
	case proto.CMSG_LOGIN:
		ret = handleLogin(s, packet)
	default:
		ret = -1
	}

	return ret
}

func handleLogin(s game.Session, p *proto.WorldPacket) int {
	log.Realm.Println("process login cmd ...")

	data := strings.Split(string(p.Data()), "/")
	log.Realm.Printf("Account:%s Pass:%s Id:%s", data[0], data[1], data[2])

	ret := 0

	packet := &proto.WorldPacket{
		Recog:  proto.NoPay,
		Opcode: proto.SMSG_LOGIN_FAILED,
	}

	s.SendPacket(packet)
	return ret
}
