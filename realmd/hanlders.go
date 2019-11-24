package realmd

import (
	"mircore/core"
	"mircore/game/proto"
	"strings"
)

func handleLogin(s core.Session, p *proto.WorldPacket) int {
	data := strings.Split(p.String(), "/")
	slog.Printf("Account:%s Pass:%s Id:%s", data[0], data[1], data[2])

	ret := 0

	packet := &proto.WorldPacket{
		Recog:  proto.WrongPwd,
		Opcode: proto.SMSG_LOGIN_FAILED,
	}

	s.SendPacket(packet)

	return ret
}

func handleRegister(s core.Session, p *proto.WorldPacket) int {
	ret := 0

	data := strings.Split(p.String(), "/")
	slog.Println(data)

	return ret
}
