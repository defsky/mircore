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
	CMSG_LOGIN        OpcodeType = 2001
	CMSG_REGISTER     OpcodeType = 2002
	CMSG_CHANGEPWD    OpcodeType = 2003
	CMSG_QUERYCHR     OpcodeType = 100 // 查询角色
	CMSG_NEWCHR       OpcodeType = 101 // 创建角色
	CMSG_DELCHR       OpcodeType = 102 // 删除角色
	CMSG_SELCHR       OpcodeType = 103 // 选择角色
	CMSG_SELECTSERVER OpcodeType = 104 // 服务器

	//SMSG_PASSWD_FAIL           OpcodeType = 503 // 验证失败,"服务器验证失败,需要重新登录"??
	SMSG_CERTIFICATION_FAIL OpcodeType = 501 //世界服务器认证失败
	SMSG_ID_NOTFOUND        OpcodeType = 502 //服务器未找到

	SMSG_LOGIN_FAILED OpcodeType = 503
	UserNotFound      RecogType  = 0
	WrongPwd          RecogType  = -1
	WrongPwd3Times    RecogType  = -2
	AlreadyLogin      RecogType  = -3
	NoPay             RecogType  = -4
	BeLock            RecogType  = -5

	SMSG_NEWID_SUCCESS         OpcodeType = 504 // 创建新账号成功
	SMSG_NEWID_FAIL            OpcodeType = 505 // 账号被占用
	SMSG_CHGPASSWD_SUCCESS     OpcodeType = 506 // 修改密码成功
	SMSG_CHGPASSWD_FAIL        OpcodeType = 507 // 修改密码失败
	SMSG_GETBACKPASSWD_SUCCESS OpcodeType = 508 // 密码找回成功
	SMSG_GETBACKPASSWD_FAIL    OpcodeType = 509 // 密码找回失败
	SMSG_QUERYCHR              OpcodeType = 520 // 返回角色信息到客户端
	SMSG_NEWCHR_SUCCESS        OpcodeType = 521 // 新建角色成功
	SMSG_NEWCHR_FAIL           OpcodeType = 522 // 新建角色失败
	SMSG_DELCHR_SUCCESS        OpcodeType = 523 // 删除角色成功
	SMSG_DELCHR_FAIL           OpcodeType = 524 // 删除角色失败
	SMSG_STARTPLAY             OpcodeType = 525 // 开始进入游戏世界(点了健康游戏忠告后进入游戏画面)
	SMSG_STARTFAIL             OpcodeType = 526 // //开始失败,玩传奇深有体会,有时选择角色,点健康游戏忠告后黑屏
	SMSG_QUERYCHR_FAIL         OpcodeType = 527 // 返回角色信息到客户端失败
	SMSG_OUTOFCONNECTION       OpcodeType = 528 // 超过最大连接数,强迫用户下线
	SMSG_PASSOK_SELECTSERVER   OpcodeType = 529 // 密码验证完成且密码正确,开始选服
	SMSG_SELECTSERVER_OK       OpcodeType = 530 // 选服成功
)

func (v OpcodeType) String() string {
	switch v {
	case CMSG_LOGIN:
		return "CMSG_LOGIN"
	case CMSG_REGISTER:
		return "CMSG_REGISTER"
	case CMSG_CHANGEPWD:
		return "CMSG_CHANGEPWD"
	case CMSG_NEWCHR:
		return "CMSG_NEWCHAR"
	case CMSG_QUERYCHR:
		return "CMSG_QUERYCHR"
	case CMSG_SELCHR:
		return "CMSG_SELCHAR"
	case CMSG_DELCHR:
		return "CMSG_DELCHAR"
	case CMSG_SELECTSERVER:
		return "CMSG_SELECTSERVER"
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
