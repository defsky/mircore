package core

import "golang.org/x/text/encoding/simplifiedchinese"

func EncodeTo(charset string, data string) string {
	switch charset {
	case "GB18030":
		buf, _ := simplifiedchinese.GB18030.NewEncoder().String(data)

		return buf
	default:
		return data
	}
}

func DecodeFrom(charset string, data []byte) string {
	switch charset {
	case "GB18030":
		buf, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(data)
		return string(buf)
	default:
		return string(data)
	}
}
