package game

import "github.com/Allenxuxu/gev/connection"

//Session world session object
type Session struct {
	conn *connection.Connection
}

//SendPacket send packet to session
func (s *Session) SendPacket(p []byte) {
	s.conn.Send(p)
}
