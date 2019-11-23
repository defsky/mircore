package proto

import (
	"bytes"
	"encoding/binary"
)

//WorldPacket packet type for world socket
type WorldPacket struct {
	Recog  RecogType
	Opcode OpcodeType
	P1     ParamType
	P2     TagType
	P3     SeriesType
	buf    *bytes.Buffer
	Seq    int
	Size   int
}

func (p *WorldPacket) Marshal() (buf *ProtoBuf) {
	var data []byte
	var dataLen int

	buf = new(ProtoBuf)
	if p.buf == nil {
		p.buf = new(bytes.Buffer)
	} else {
		dataLen = p.buf.Len()
		if dataLen > 0 {
			data = p.buf.Bytes()
			p.buf.Reset()
		}
	}

	binary.Write(buf, binary.LittleEndian, p.Recog)
	binary.Write(buf, binary.LittleEndian, p.Opcode)
	binary.Write(buf, binary.LittleEndian, p.P1)
	binary.Write(buf, binary.LittleEndian, p.P2)
	binary.Write(buf, binary.LittleEndian, p.P3)

	if data != nil {
		buf.Write(data)
	}

	return
}

func (p *WorldPacket) SetRecogOp(n RecogType, m OpcodeType) {

}

func (p *WorldPacket) SetParams(p1 ParamType, p2 TagType, p3 SeriesType) {

}

func (p *WorldPacket) Data() []byte {
	return p.buf.Bytes()
}
