package protocol

import (
	"encoding/binary"
	"errors"
	"strconv"

	"github.com/Allenxuxu/gev/connection"
	"github.com/Allenxuxu/ringbuffer"
	"github.com/gobwas/pool/pbytes"

	"mircore/utils"
	"mircore/utils/log"
)

//GameProtocol game protocol struct
type GameProtocol struct {
}

//PacketHeader game packet header struct
type PacketHeader struct {
	Seq        int
	Opcode     uint16
	PacketSize int
	Recog      uint32
	P1         uint16
	P2         uint16
	P3         uint16
}

//UnPacket unpack data
func (p *GameProtocol) UnPacket(c *connection.Connection, buffer *ringbuffer.RingBuffer) (interface{}, []byte) {
	buf := pbytes.GetLen(1)
	defer pbytes.Put(buf)

	data := make([]byte, 0, 1)
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

			dataPacket, err := decodeData(data)
			if err != nil {
				return nil, nil
			}
			log.Core.Println("Received:", utils.RawData(dataPacket))

			header, body := unPackHeader(dataPacket)
			header.Seq = packetSeq

			return header, body
		}
		if isData {
			data = append(data, buf...)
		}
	}

	buffer.VirtualRevert()
	return nil, nil
}

//Packet pack data
func (p *GameProtocol) Packet(c *connection.Connection, data []byte) []byte {
	return data
}

func unPackHeader(packet []byte) (header *PacketHeader, data []byte) {
	header = &PacketHeader{}

	header.Recog = binary.LittleEndian.Uint32(packet[0:4])
	header.Opcode = binary.LittleEndian.Uint16(packet[4:6])
	header.P1 = binary.LittleEndian.Uint16(packet[6:8])
	header.P2 = binary.LittleEndian.Uint16(packet[8:10])
	header.P3 = binary.LittleEndian.Uint16(packet[10:12])
	header.PacketSize = len(packet)

	if header.PacketSize > 12 {
		data = make([]byte, header.PacketSize-12)
		copy(data, packet[12:])
	}

	return
}

func decodeData(s []byte) (d []byte, err error) {
	srcLen := len(s)

	buf := pbytes.GetLen(4)
	defer pbytes.Put(buf)

	for i, v := range s {
		v = v - 0x3c
		if v < 0 {
			return nil, errors.New("Invalid Packet")
		}
		s[i] = v & 0x3f
	}

	data := make([]byte, 0, 1)
	dataLen := 0
	i, j := 0, 4
	for {
		if j > srcLen {
			j = srcLen
		}
		dataLen = (j - i) * 3 / 4
		if dataLen < 1 {
			break
		}
		copy(buf, s[i:j])

		var tmpUint32 uint32 = 0

		for i, v := range buf {
			n := uint((3 - i) * 6)
			tmpUint32 = tmpUint32 | uint32(v)<<n
		}

		binary.BigEndian.PutUint32(buf, tmpUint32)

		if dataLen > 0 {
			data = append(data, buf[1:1+dataLen]...)
		}

		i, j = i+4, j+4

		if i >= srcLen {
			break
		}
	}

	return data, nil
}
func encodeData(s []byte) []byte {
	return s
}
