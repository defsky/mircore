package realm

import (
	"container/list"
	"strconv"
	"sync"

	"github.com/Allenxuxu/gev"
	"github.com/Allenxuxu/gev/connection"

	"mircore/protocol"
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
	corelog.Realm.Println(v...)
}

func logf(fmt string, v ...interface{}) {
	corelog.Realm.Printf(fmt, v...)
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

	c.Send([]byte("fvpvTbKVC\\WnpqQvh_xdY\\\\"))
}

//OnMessage callback when new message come
func (s *LoginServer) OnMessage(c *connection.Connection, ctx interface{}, data []byte) (out []byte) {
	header := ctx.(*protocol.PacketHeader)

	logf("Seq(%d) Recog(%d) Opcode(%d) PacketSize(%d): %s", 
		header.Seq, header.Recog, header.Opcode, header.PacketSize,string(data))

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
