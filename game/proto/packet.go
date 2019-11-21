package proto

import "encoding/binary"

//WorldPacketHeader game packet header struct
type WorldPacketHeader struct {
	Seq    int
	Opcode uint16
	Recog  uint32
	P1     uint16
	P2     uint16
	P3     uint16
}

//WorldPacket packet type for world socket
type WorldPacket struct {
	Header *WorldPacketHeader
	Data   []byte
	Size   int
}

//NewPacket create a WorldPacket
func NewPacket(data []byte) (p *WorldPacket) {
	p = new(WorldPacket)
	p.Header = new(WorldPacketHeader)

	p.Size = len(data)

	headerSize := 12

	if data[0] == '*' && data[1] == '*' {
		headerSize = 2
	}

	if headerSize == 2 {
		p.Header.Opcode = 65001
	} else {
		p.Header.Recog = binary.LittleEndian.Uint32(data[0:4])
		p.Header.Opcode = binary.LittleEndian.Uint16(data[4:6])
		p.Header.P1 = binary.LittleEndian.Uint16(data[6:8])
		p.Header.P2 = binary.LittleEndian.Uint16(data[8:10])
		p.Header.P3 = binary.LittleEndian.Uint16(data[10:12])
	}

	dataSize := p.Size - headerSize

	if dataSize > 0 {
		p.Data = make([]byte, dataSize)
		copy(p.Data, data[headerSize:])
	}

	return
}

//GetOpcode get opcode of current packet
func (p *WorldPacket) GetOpcode() uint16 {
	return p.Header.Opcode
}
