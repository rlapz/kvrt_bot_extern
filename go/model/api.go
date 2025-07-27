package model

type ApiArgs struct {
	CmdName   string
	ChatId    int64
	UserId    int64
	MessageId int64
	Api       string
	Config    string
	Text      string
	Data      string
}

type ApiReq struct {
	Type      string `json:"type"`
	ChatId    int64  `json:"chat_id"`
	UserId    int64  `json:"user_id"`
	MessageId int64  `json:"message_id"`
	Deletable bool   `json:"deletable"`
	Text      string `json:"text"`
}
