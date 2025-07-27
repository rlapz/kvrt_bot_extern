package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/rlapz/kvrt_bot_extern/extra"
	"github.com/rlapz/kvrt_bot_extern/model"
)

func main() {
	// TODO: verify argument list

	chat_id, err := strconv.ParseInt(os.Args[3], 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	user_id, err := strconv.ParseInt(os.Args[4], 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	message_id, err := strconv.ParseInt(os.Args[5], 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	req := model.ApiArgs{
		CmdName:   os.Args[1],
		ChatId:    chat_id,
		UserId:    user_id,
		MessageId: message_id,
		Api:       os.Getenv("TG_API"),
		Config:    os.Getenv("TG_CONFIG_FILE"),
		Text:      os.Args[6],
		Data:      os.Args[7],
	}

	switch os.Args[1] {
	case "/neko2":
		extra.RunNeko(&req)
	default:
		fmt.Println("well, nice try")
	}
}
