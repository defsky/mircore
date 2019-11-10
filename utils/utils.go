package utils

import (
	"fmt"
)

//RawData returns data package in decimal form
func RawData(data []byte) string {
	var ret string = "\n    ["

	for i, v := range data {
		if i > 0 {
			if i%16 == 0 {
				ret += "]\n    ["
			} else if i%8 == 0 {
				ret += "| "
			} else {
				ret += " "
			}
		}
		ret += fmt.Sprintf("%-3d", int(v))
	}
	ret += "]"

	return ret
}
