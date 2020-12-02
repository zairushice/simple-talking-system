package message

const (
	LoginMsgType       = "LoginMsg"
	LoginResMsgType    = "LoginResMsg"
	RegisterMsgType    = "RegisterMsg"
	RegisterResMsgType = "RegisterResMsg"
)

type RegisterMsg struct {
	User User `json:"user"`
}

type RegisterResMsg struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

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
