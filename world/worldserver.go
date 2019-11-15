package world

import (
	"container/list"
	"strconv"
	"sync"

	"github.com/Allenxuxu/gev"
	"github.com/Allenxuxu/gev/connection"

	"mircore/protocol"
	"mircore/utils"
	corelog "mircore/utils/log"
)

func log(v ...interface{}) {
	corelog.World.Println(v...)
}

func logf(fmt string, v ...interface{}) {
	corelog.World.Printf(fmt, v...)
}

//WorldServer world server
type WorldServer struct {
	conns    *list.List
	mtx      sync.RWMutex
	port     int
	listener *gev.Server
}

//NewWorldServer create new world server
func NewWorldServer(port int) (*WorldServer, error) {
	var err error

	s := new(WorldServer)
	s.port = port
	s.conns = list.New()

	s.listener, err = gev.NewServer(s,
		gev.Network("tcp"),
		gev.Address(":"+strconv.Itoa(port)),
		gev.NumLoops(-1),
		gev.Protocol(&protocol.GameProtocol{}))

	return s, err
}

//OnConnect new connection callback
func (s *WorldServer) OnConnect(c *connection.Connection) {
	logf("New Connection from: %s\n", c.PeerAddr())

	s.mtx.Lock()
	e := s.conns.PushBack(c)
	s.mtx.Unlock()
	c.SetContext(e)
}

//OnMessage new message callback
func (s *WorldServer) OnMessage(c *connection.Connection, ctx interface{}, data []byte) (out []byte) {
	log("Received: ", utils.RawData(data))

	return
}

//OnClose connection close callback
func (s *WorldServer) OnClose(c *connection.Connection) {
	logf("Connection closed from: %s", c.PeerAddr())

	e := c.Context().(*list.Element)

	s.mtx.Lock()
	s.conns.Remove(e)
	s.mtx.Unlock()
}

//Start start world server
func (s *WorldServer) Start() {
	logf("Listening at *:%d\n", s.port)

	s.listener.Start()
}

//Stop stop world server
func (s *WorldServer) Stop() {
	log("Stopping server ...")
	s.listener.Stop()
}
