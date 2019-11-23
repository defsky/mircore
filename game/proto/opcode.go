package proto

//RecogType recog data type
type RecogType int32

//OpcodeType opcode data type
type OpcodeType uint16

//ParamType param data type
type ParamType uint16

//TagType type data type
type TagType uint16

//SeriesType series data type
type SeriesType uint16

const (
	CMSG_LOGIN     OpcodeType = 2001
	CMSG_REGISTER  OpcodeType = 2002
	CMSG_CHANGEPWD OpcodeType = 2003

	SMSG_LOGIN_FAILED OpcodeType = 503
	UserNotFound      RecogType  = 0
	WrongPwd          RecogType  = -1
	WrongPwd3Times    RecogType  = -2
	AlreadyLogin      RecogType  = -3
	NoPay             RecogType  = -4
	BeLock            RecogType  = -5
)

func (v OpcodeType) String() string {
	switch v {
	case CMSG_LOGIN:
		return "CMSG_LOGIN"
	case CMSG_REGISTER:
		return "CMSG_REGISTER"
	case CMSG_CHANGEPWD:
		return "CMSG_CHANGEPWD"
	case SMSG_LOGIN_FAILED:
		return "SMSG_LOGIN_FAILED"
	default:
		return "UNKOWN_OPCODE"
	}
}
func (v RecogType) String() string {
	switch v {
	case UserNotFound:
		return "UserNotFound"
	case WrongPwd:
		return "WrongPwd"
	case WrongPwd3Times:
		return "WrongPwd3Times"
	case AlreadyLogin:
		return "AlreadyLogin"
	case NoPay:
		return "NoPay"
	case BeLock:
		return "BeLock"
	default:
		return "UNKOWN_RECOG"
	}
}
