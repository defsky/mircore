package proto

import (
	"fmt"
	"mircore/utils"
	"testing"
)

func TestEncodeData(t *testing.T) {
	buf := []byte{0x0, 0x0, 0x0, 0x0, 0xF7, 0x01, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}

	buf2 := encodeData(buf)
	fmt.Printf("%s\n", utils.RawData(buf2))

	buf3, err := decodeData(buf2)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%s\n", utils.RawData(buf3))
}
