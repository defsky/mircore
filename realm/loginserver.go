package realm

import (
	"container/list"
	"strconv"
	"sync"

	"github.com/Allenxuxu/gev"
	"github.com/Allenxuxu/gev/connection"
	_ "github.com/Allenxuxu/gev/log"

	"mircore/protocol"
	"mircore/utils"
	corelog "mircore/utils/log"
)

//LoginServer login server struct
type LoginServer struct {
	conns *list.List
	mtx   sync.RWMutex
	port  int

	listener *gev.Server
}

func log(v ...interface{}) {
	corelog.RealmLog.Println(v...)
}

func logf(fmt string, v ...interface{}) {
	corelog.RealmLog.Printf(fmt, v...)
}

//NewLoginServer create login server
func NewLoginServer(port int) (*LoginServer, error) {
	var err error

	s := new(LoginServer)
	s.port = port
	s.conns = list.New()

	s.listener, err = gev.NewServer(s,
		gev.Network("tcp"),
		gev.Address(":"+strconv.Itoa(port)),
		gev.NumLoops(-1),
		gev.Protocol(&protocol.GameProtocol{}))

	return s, err
}

//OnConnect callback when new connection come
func (s *LoginServer) OnConnect(c *connection.Connection) {
	logf("New Connection from: %s\n", c.PeerAddr())

	s.mtx.Lock()
	e := s.conns.PushBack(c)
	s.mtx.Unlock()
	c.SetContext(e)
}

//OnMessage callback when new message come
func (s *LoginServer) OnMessage(c *connection.Connection, ctx interface{}, data []byte) (out []byte) {
	log("Received: ", utils.RawData(data))

	return
}

//OnClose callback when connection close
func (s *LoginServer) OnClose(c *connection.Connection) {
	logf("Connection closed from: %s", c.PeerAddr())

	e := c.Context().(*list.Element)

	s.mtx.Lock()
	s.conns.Remove(e)
	s.mtx.Unlock()
}

//Start start server
func (s *LoginServer) Start() {
	logf("Listening at *:%d\n", s.port)

	s.listener.Start()
}

//Stop stop server
func (s *LoginServer) Stop() {
	log("Stopping server ...")
	s.listener.Stop()
}
