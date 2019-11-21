package realmd

import (
	"encoding/binary"
	"mircore/game/proto"
	"mircore/utils/log"
	"strings"
)

//handleIncoming process incoming packet
func handleIncoming(packet *proto.WorldPacket) (out []byte, ret int) {
	opcode := packet.GetOpcode()

	switch opcode {
	case proto.CmsgLogin:
		out, ret = handleLogin(packet)
	default:
		ret = -1
	}

	return
}

func handleLogin(p *proto.WorldPacket) (out []byte, ret int) {
	log.Realm.Println("process login cmd ...")

	data := strings.Split(string(p.Data), "/")
	log.Realm.Printf("Account:%s Pass:%s Id:%s", data[0], data[1], data[2])

	ret = 0

	out = make([]byte, 0, 12)

	recog := make([]byte, 4)
	binary.LittleEndian.PutUint32(recog, 0)
	out = append(out, recog...)

	opcode := make([]byte, 2)
	binary.LittleEndian.PutUint16(opcode, proto.SmsgAuthFail)
	out = append(out, opcode...)

	out = out[:12]

	return
}
