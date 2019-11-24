package core

import "mircore/game/proto"

type Session interface {
	SendPacket(p *proto.WorldPacket)
}
