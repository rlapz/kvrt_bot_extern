package model

type TgUpdate struct {
	UpdateId int64     `json:"update_id"`
	Message  TgMessage `json:"message"`
}

type TgMessage struct {
	Id      int64      `json:"message_id"`
	ReplyTo *TgMessage `json:"reply_to_message"`
	Text    string     `json:"text"`
}
