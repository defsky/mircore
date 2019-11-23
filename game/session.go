package game

import (
	"mircore/game/proto"
	"mircore/utils/log"

	"github.com/Allenxuxu/gev/connection"
)

type Session interface {
	SendPacket(p *proto.WorldPacket)
}

//Session world session object
type RealmSession struct {
	Conn *connection.Connection
}

//SendPacket send packet to session
func (s *RealmSession) SendPacket(p *proto.WorldPacket) {
	pbuf := p.Marshal()
	log.Realm.Printf("Send:%s\n", pbuf.HexString())

	s.Conn.Send(pbuf.Encode().Bytes())
}

type WorldSession struct {
	Conn *connection.Connection
}
