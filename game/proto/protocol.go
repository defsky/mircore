package proto

import (
	"encoding/binary"
	"errors"
	"math"
	"mircore/utils"
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

			packet := NewPacket(dataPacket)
			packet.Header.Seq = packetSeq

			return packet, nil
		}
		if isData {
			data = append(data, buf...)
		}
	}

	buffer.VirtualRevert()
	return nil, nil
}

//Packet pack data
func (p *Protocol) Packet(c *connection.Connection, data []byte) []byte {
	out := encodeData(data)
	log.Core.Printf("Send: %s", string(out))
	log.Core.Printf("%s\n", utils.RawData(out))

	return out
}

func decodeData(in []byte) (out []byte, err error) {
	buf := pbytes.GetLen(4)
	defer pbytes.Put(buf)

	srcLen := len(in)

	for i, v := range in {
		v = v - 0x3c
		if v < 0 {
			return nil, errors.New("Invalid Packet")
		}
		in[i] = v & 0x3f
	}

	out = make([]byte, 0, srcLen*3/4)
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
		copy(buf, in[i:j])

		var tmpUint32 uint32 = 0

		for i, v := range buf {
			n := uint((3 - i) * 6)
			tmpUint32 = tmpUint32 | uint32(v)<<n
		}

		binary.BigEndian.PutUint32(buf, tmpUint32)

		out = append(out, buf[1:1+dataLen]...)

		i, j = i+4, j+4

		if i >= srcLen {
			break
		}
	}

	return
}
func encodeData(in []byte) (out []byte) {
	srcLen := len(in)
	outLen := int(math.Ceil(float64(srcLen*4/3))) + 2
	out = make([]byte, outLen)

	out[0] = '#'

	dataLen := 4
	i, j := 0, 3
	idx := 1
	for {
		if j > srcLen {
			j = srcLen
			dataLen = int(math.Ceil(float64((j - i) * 4 / 3)))
		}

		buf := make([]byte, 4)
		copy(buf[1:], in[i:j])

		tempUint32 := binary.BigEndian.Uint32(buf)
		for a := 0; a < dataLen; a++ {
			n := uint((3 - a) * 6)
			out[idx] = byte((tempUint32>>n)&0x3f + 0x3c)
			idx++
		}

		i, j = i+3, j+3

		if i >= srcLen {
			break
		}
	}

	out[outLen-1] = '!'

	return
}
