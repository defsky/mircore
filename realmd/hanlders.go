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
	infos := parseRegInfo(p)

	slog.Println("RegInfo:", infos)

	packet := &proto.WorldPacket{
		Recog:  0,
		Opcode: proto.SMSG_NEWID_FAIL,
	}

	s.SendPacket(packet)

	return ret
}

func parseRegInfo(p *proto.WorldPacket) (info RegInfo) {
	var dataLen int
	info = make(RegInfo)

	for _, v := range regInfoFields {
		for k, v := range v {
			buf := make([]byte, v+1)

			n, _ := p.Read(buf)
			if n != v+1 {
				return
			}

			dataLen = int(buf[0])
			if dataLen != 0 {
				info[k] = core.DecodeFrom("GB18030", buf[1:dataLen+1])//string(buf[1 : dataLen+1])
			} else {
				info[k] = ""
			}
		}
	}

	return
}
