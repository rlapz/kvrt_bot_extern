package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/rlapz/kvrt_bot_extern/extra"
	"github.com/rlapz/kvrt_bot_extern/model"
)

const (
	_ARG_EXEC_FILE = iota
	_ARG_CMD_NAME
	_ARG_EXEC_TYPE
	_ARG_CHAT_FLAGS
	_ARG_CHAT_ID
	_ARG_USER_ID
	_ARG_MESSAGE_ID
	_ARG_CHAT_TEXT
	_ARG_RAW_JSON
)

func runCmd(r *model.ApiArgs) {
	switch r.CmdName {
	case "/neko":
		extra.RunNeko(r)
	default:
		fmt.Println("well, nice try!")
	}
}

func main() {
	// TODO: verify argument list
	fmt.Println("args:", os.Args)

	chat_flags, err := strconv.ParseInt(os.Args[_ARG_CHAT_FLAGS], 10, 32)
	if err != nil {
		fmt.Println(err)
		return
	}

	chat_id, err := strconv.ParseInt(os.Args[_ARG_CHAT_ID], 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	user_id, err := strconv.ParseInt(os.Args[_ARG_USER_ID], 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	message_id, err := strconv.ParseInt(os.Args[_ARG_MESSAGE_ID], 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	owner_id, err := strconv.ParseInt(os.Getenv("TG_OWNER_ID"), 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	bot_id, err := strconv.ParseInt(os.Getenv("TG_BOT_ID"), 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	req := model.ApiArgs{
		CmdName:     os.Args[_ARG_CMD_NAME],
		ChatFlags:   int(chat_flags),
		ChatId:      chat_id,
		UserId:      user_id,
		MessageId:   message_id,
		Api:         os.Getenv("TG_API"),
		RootDir:     os.Getenv("TG_ROOT_DIR"),
		ConfigFile:  os.Getenv("TG_CONFIG_FILE"),
		TgApi:       os.Getenv("TG_API_URL"),
		OwnerId:     owner_id,
		BotId:       bot_id,
		BotUsername: os.Getenv("TG_BOT_USERNAME"),
		Text:        os.Args[_ARG_CHAT_TEXT],
		Data:        os.Args[_ARG_RAW_JSON],
	}

	runCmd(&req)
}
