package core

import (
	"container/list"
	"mircore/game"
	"mircore/utils/log"

	"strconv"
	"sync"

	"github.com/Allenxuxu/gev"
	"github.com/Allenxuxu/gev/connection"
)

//Realm login server struct
type Realm struct {
	conns *list.List
	mtx   sync.RWMutex
	port  int

	listener *gev.Server
}

//NewRealm create login server
func NewRealm(port int, proto connection.Protocol) (*Realm, error) {
	var err error

	s := new(Realm)
	s.port = port
	s.conns = list.New()

	s.listener, err = gev.NewServer(s,
		gev.Network("tcp"),
		gev.Address(":"+strconv.Itoa(port)),
		gev.NumLoops(-1),
		gev.Protocol(proto))

	return s, err
}

//OnConnect callback when new connection come
func (s *Realm) OnConnect(c *connection.Connection) {
	log.Realm.Printf("New Connection from: %s\n", c.PeerAddr())

	s.mtx.Lock()
	e := s.conns.PushBack(c)
	s.mtx.Unlock()
	c.SetContext(e)

	c.Send([]byte("fvpvTbKVC\\WnpqQvh_xdY\\\\"))
}

//OnMessage callback when new message come
func (s *Realm) OnMessage(c *connection.Connection, ctx interface{}, data []byte) (out []byte) {
	header := ctx.(*game.PacketHeader)

	log.Realm.Printf("Seq(%d) Recog(%d) Opcode(%d) PacketSize(%d): %s",
		header.Seq, header.Recog, header.Opcode, header.PacketSize, string(data))

	return
}

//OnClose callback when connection close
func (s *Realm) OnClose(c *connection.Connection) {
	log.Realm.Printf("Connection closed from: %s", c.PeerAddr())

	e := c.Context().(*list.Element)

	s.mtx.Lock()
	s.conns.Remove(e)
	s.mtx.Unlock()
}

//Start start server
func (s *Realm) Start() {
	log.Realm.Printf("Listening at *:%d\n", s.port)

	s.listener.Start()
}

//Stop stop server
func (s *Realm) Stop() {
	log.Realm.Println("Stopping server ...")
	s.listener.Stop()
}