package tg

type Update struct {
	UpdateId int64   `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Id      int64    `json:"message_id"`
	ReplyTo *Message `json:"reply_to_message"`
	Text    string   `json:"text"`
}
