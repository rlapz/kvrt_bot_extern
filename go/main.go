package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/rlapz/kvrt_bot_extern/api"
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

func registerCmd() map[string]func(*model.ApiArgs) {
	return map[string]func(*model.ApiArgs){
		"/neko":     extra.RunNeko,
		"/waifu":    extra.RunWaifu,
		"/aniquote": extra.RunAniquote,
		"/s":        extra.RunSed,
		"/tellme":   extra.RunTellMe,
		"/joke":     extra.RunJoke,
		"/darkjoke": extra.RunDarkJoke,
		"/advice":   extra.RunAdvice,
		"/stoicism": extra.RunStoicism,
		"/quote":    extra.RunQuote,
	}
}

func runCmd(r *model.ApiArgs) {
	cmdMap := registerCmd()
	handler, ok := cmdMap[r.CmdName]
	if !ok {
		fmt.Println("well, nice try!")
		return
	}

	req := model.ApiReq{
		Type:    "acquire",
		ChatId:  r.ChatId,
		UserId:  r.UserId,
		Context: r.CmdName,
	}

	err := api.Submit(r, "session", &req)
	if err != nil {
		_ = api.SendTextPlain(r, "Please wait!")
		return
	}

	defer func() {
		req.Type = "release"
		err := api.Submit(r, "session", &req)
		if err != nil {
			fmt.Println("error:", err)
		}
	}()

	handler(r)
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
		DbFile:      os.Getenv("TG_DB_FILE"),
		TgApi:       os.Getenv("TG_API_URL"),
		OwnerId:     owner_id,
		BotId:       bot_id,
		BotUsername: os.Getenv("TG_BOT_USERNAME"),
		Text:        os.Args[_ARG_CHAT_TEXT],
		RawJSON:     os.Args[_ARG_RAW_JSON],
	}

	runCmd(&req)
}
