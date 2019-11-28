package proto

import (
	"mircore/utils/log"
	"strconv"

	"github.com/Allenxuxu/gev/connection"
	"github.com/Allenxuxu/ringbuffer"
	"github.com/gobwas/pool/pbytes"
)

//Protocol game protocol struct
type Protocol struct {
}

//UnPacket unpack data
func (p *Protocol) UnPacket(c *connection.Connection, buffer *ringbuffer.RingBuffer) (interface{}, []byte) {
	buf := pbytes.GetLen(1)
	defer pbytes.Put(buf)

	//data := make([]byte, 0, 1)
	pbuf := new(ProtoBuf)

	isData := false
	packetSeq := 0

	for n, _ := buffer.VirtualRead(buf); n > 0; n, _ = buffer.VirtualRead(buf) {
		if buf[0] == '#' {
			buffer.VirtualRead(buf)
			packetSeq, _ = strconv.Atoi(string(buf[0]))
			isData = true
			continue
		}
		if buf[0] == '!' {
			buffer.VirtualFlush()

			dataPacket, err := pbuf.Decode() // decodeData(data)
			if err != nil {
				return nil, nil
			}
			log.Core.Println("Received:", dataPacket.HexString())
			dataSize := dataPacket.Len()

			packet, err := dataPacket.UnMarshal()
			if err != nil {
				log.Core.Println("invalid packet, ignore", err)
				return nil, nil
			}

			packet.Seq = packetSeq
			packet.Size = dataSize

			return packet, nil
		}
		if isData {
			//data = append(data, buf...)
			pbuf.Write(buf)
		}
	}

	buffer.VirtualRevert()
	return nil, nil
}

//Packet pack data
func (p *Protocol) Packet(c *connection.Connection, data []byte) []byte {
	log.Core.Printf("Send: %s", string(data))

	return data
}
