package realmd

import (
	"container/list"
	"mircore/game/proto"
	"mircore/utils/log"

	"strconv"
	"sync"
	"time"

	"github.com/Allenxuxu/gev"
	"github.com/Allenxuxu/gev/connection"
)

var slog = log.Realm

//Realm login server struct
type Realm struct {
	conns *list.List
	mtx   sync.RWMutex
	port  int

	listener *gev.Server
}

//NewRealmServer create login server
func NewRealmServer(port int, proto connection.Protocol) (*Realm, error) {
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

func (s *Realm) report() {
	s.mtx.Lock()
	ConnCounts := s.conns.Len()
	s.mtx.Unlock()
	slog.Printf("Alive Connections: %d\n", ConnCounts)
}

//OnConnect callback when new connection come
func (s *Realm) OnConnect(c *connection.Connection) {
	slog.Printf("New Connection from: %s\n", c.PeerAddr())

	sess := &RealmSession{
		Conn:      c,
		StartTime: time.Now(),
	}

	s.mtx.Lock()
	e := s.conns.PushBack(sess)
	s.mtx.Unlock()
	c.SetContext(e)

	c.Send([]byte("fvpvTbKVC\\WnpqQvh_xdY\\\\"))
}

//OnMessage callback when new message come
func (s *Realm) OnMessage(c *connection.Connection, ctx interface{}, data []byte) (out []byte) {
	packet := ctx.(*proto.WorldPacket)

	slog.Printf("Seq(%d) %s(%d) Recog(%d) Size(%d): %s",
		packet.Seq, packet.Opcode, packet.Opcode, packet.Recog,
		packet.Size, packet)

	e := c.Context().(*list.Element)
	sess := e.Value.(*RealmSession)
	sess.HandlePacket(packet)

	return
}

//OnClose callback when connection close
func (s *Realm) OnClose(c *connection.Connection) {
	slog.Printf("Connection closed from: %s", c.PeerAddr())

	e := c.Context().(*list.Element)

	s.mtx.Lock()
	s.conns.Remove(e)
	s.mtx.Unlock()
}

//Start start server
func (s *Realm) Start() {
	slog.Printf("Listening at *:%d\n", s.port)

	s.listener.RunEvery(time.Second*10, s.report)

	s.listener.Start()
}

//Stop stop server
func (s *Realm) Stop() {
	slog.Println("Stopping server ...")
	s.listener.Stop()
}
