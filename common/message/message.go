package message

const (
	LoginMsgType       = "LoginMsg"
	LoginResMsgType    = "LoginResMsg"
	RegisterMsgType    = "RegisterMsg"
	RegisterResMsgType = "RegisterResMsg"
)

type RegisterMsg struct {
	UserId       int    `json:"userId"`
	UserPassword string `json:"password"`
	Email        string `json:"email"`
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
