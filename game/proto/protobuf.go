package proto

import (
	"bytes"
	"encoding/binary"
	"errors"
	"math"
	"mircore/utils"
	"mircore/utils/log"

	"github.com/Allenxuxu/gev/connection"
	"github.com/gobwas/pool/pbytes"
)

type ProtoBuf struct {
	bytes.Buffer
}

func (p *ProtoBuf) Encode() *ProtoBuf {
	buf := pbytes.GetLen(p.Len())
	defer pbytes.Put(buf)

	p.Read(buf)
	p.Reset()

	p.Write(encodeData(buf))

	return p
}

func (p *ProtoBuf) Decode() (*ProtoBuf, error) {
	buf := pbytes.GetLen(p.Len())
	defer pbytes.Put(buf)

	p.Read(buf)
	p.Reset()

	data, err := decodeData(buf)
	if err != nil {
		return p, err
	}

	p.Write(data)

	return p, nil
}

func (p *ProtoBuf) UnMarshal() (*WorldPacket, error) {
	return unMarshal(p.Bytes())
}
func (p *ProtoBuf) HexString() string {
	return utils.RawData(p.Bytes())
}

func (p *ProtoBuf) String() string {
	return string(p.Bytes())
}

func unMarshal(data []byte) (p *WorldPacket, err error) {
	p = &WorldPacket{
		buf: bytes.NewBuffer(data),
	}

	err = binary.Read(p.buf, binary.LittleEndian, &p.Recog)
	if err != nil {
		return nil, err
	}

	err = binary.Read(p.buf, binary.LittleEndian, &p.Opcode)
	if err != nil {
		return nil, err
	}

	err = binary.Read(p.buf, binary.LittleEndian, &p.P1)
	if err != nil {
		return nil, err
	}
	err = binary.Read(p.buf, binary.LittleEndian, &p.P2)
	if err != nil {
		return nil, err
	}
	err = binary.Read(p.buf, binary.LittleEndian, &p.P2)
	if err != nil {
		return nil, err
	}

	return
}

//Packet pack data
func (p *Protocol) Packet(c *connection.Connection, data []byte) []byte {
	log.Core.Printf("Send: %s", string(data))

	return data
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
