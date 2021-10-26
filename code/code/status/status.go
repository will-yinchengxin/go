package status

type ResCode int64

const (
	CodeSucces ResCode = iota + 1001
	CodeFail
	CodeUnLogin
	CodeServerBusy
)

var codeMsg = map[ResCode]string{
	CodeSucces:     "this is success status",
	CodeFail:       "this is fail status",
	CodeUnLogin:    "this is unlogin status",
	CodeServerBusy: "the server is busy",
}

func (code ResCode) Msg() string {
	msg, ok := codeMsg[code]
	if !ok {
		msg = codeMsg[CodeServerBusy]
	}
	return msg
}
