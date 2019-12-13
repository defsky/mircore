package realmd

import (
	"mircore/game/proto"
	"time"

	"github.com/Allenxuxu/gev/connection"
)

//Session world session object
type RealmSession struct {
	Conn      *connection.Connection
	StartTime time.Time
}

//SendPacket send packet to session
func (s *RealmSession) SendPacket(p *proto.WorldPacket) {
	pbuf := p.Marshal()
	slog.Printf("Send:%s\n", pbuf.HexString())

	s.Conn.Send(pbuf.Encode().Bytes())
}

func (s *RealmSession) HandlePacket(p *proto.WorldPacket) {
	ret := 0

	slog.Printf("handle %s(%d) from %s", p.Opcode, p.Opcode, s.Conn.PeerAddr())

	switch p.Opcode {
	case proto.CMSG_LOGIN:
		ret = handleLogin(s, p)
	case proto.CMSG_REGISTER:
		ret = handleRegister(s, p)
	default:
		ret = -1
		slog.Printf("Unkown Opcode %d, ignore...", p.Opcode)
	}

	if ret < 0 {

	}
}
