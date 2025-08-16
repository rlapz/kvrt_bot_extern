package model

const (
	CHAT_FLAG_ALLOW_CMD_NSFW   = (1 << 0)
	CHAT_FLAG_ALLOW_CMD_EXTERN = (1 << 1)
	CHAT_FLAG_ALLOW_CMD_EXTRA  = (1 << 2)
)

type ApiArgs struct {
	CmdName     string
	ChatFlags   int
	ChatId      int64
	UserId      int64
	MessageId   int64
	Api         string
	RootDir     string
	ConfigFile  string
	TgApi       string
	OwnerId     int64
	BotId       int64
	BotUsername string
	Text        string
	RawJSON     string
}

type ApiReq struct {
	Type      string `json:"type"`
	ChatId    int64  `json:"chat_id"`
	UserId    int64  `json:"user_id"`
	MessageId int64  `json:"message_id"`
	Text      string `json:"text"`
	TextType  string `json:"text_type"`
	Photo     string `json:"photo"`
}
