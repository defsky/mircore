package protocol

import (
	"github.com/Allenxuxu/gev/connection"
	"github.com/Allenxuxu/ringbuffer"
)

type GameProtocol struct {
}

func (p *GameProtocol) UnPacket(c *connection.Connection, buffer *ringbuffer.RingBuffer) (interface{}, []byte) {
	ret := buffer.Bytes()
	buffer.RetrieveAll()

	return nil, ret
}

func (p *GameProtocol) Packet(c *connection.Connection, data []byte) []byte {
	return data
}
