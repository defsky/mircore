package game

import "mircore/core"

//Player game player object
type Player struct {
	session *core.Session
}

//GetSession get session of player
func (p *Player) GetSession() *core.Session {
	return p.session
}
