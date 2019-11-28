package realmd

import (
	"fmt"
	"mircore/game/proto"
)

// type RegInfo struct {
// 	account     string // 10 bytes
// 	password    string // 10 bytes
// 	name        string // 20 bytes
// 	idcard      string // 14 bytes
// 	phone       string // 14 bytes
// 	question1   string // 20 bytes
// 	answer1     string // 12 bytes
// 	email       string // 51 bytes
// 	question2   string // 20 bytes
// 	answer2     string // 12 bytes
// 	birthday    string // 10 bytes
// 	mobilephone string // 55 bytes
// }

type FieldInfo map[string]int
type RegInfo map[string]string

var regInfoFields = []FieldInfo{
	{"username": 10},
	{"password": 10},
	{"name": 20},
	{"idcard": 14},
	{"phone": 14},
	{"question1": 20},
	{"answer1": 12},
	{"email": 51},
	{"question2": 20},
	{"answer2": 12},
	{"birthday": 10},
	{"mobilephone": 55},
}

func (r RegInfo) String() string {
	buf := "\n"

	for k, v := range r {
		buf += fmt.Sprintf("    [%-12s: %-20s]\n", k, v)
	}

	return buf
}

type LoginFailReason proto.RecogType
