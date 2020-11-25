package message

const (
	LoginMsgType    = "LoginMes"
	LoginResMsgType = "LoginResMsg"
)

type LoginMsg struct {
	UserId       int    `json:"userId"`
	UserPassword string `json:"password"`
	UserName     string `json:"userName"`
}

type LoginResMsg struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}
