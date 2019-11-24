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
	bytes.Buffer
	Seq  int
	Size int
}

//Marshal 把WorldPacket对象序列化为ProtoBuf
func (p *WorldPacket) Marshal() (buf *ProtoBuf) {
	buf = new(ProtoBuf)

	dataLen := p.Len()

	binary.Write(buf, binary.LittleEndian, p.Recog)
	binary.Write(buf, binary.LittleEndian, p.Opcode)
	binary.Write(buf, binary.LittleEndian, p.P1)
	binary.Write(buf, binary.LittleEndian, p.P2)
	binary.Write(buf, binary.LittleEndian, p.P3)

	if dataLen > 0 {
		buf.Write(p.Bytes())
	}

	return
}

func (p *WorldPacket) String() string {
	return string(p.Bytes())
}
