package game

//Player game player object
type Player struct {
	session *Session
}

//GetSession get session of player
func (p *Player) GetSession() *Session {
	return p.session
}
