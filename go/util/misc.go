package util

import (
	"encoding/json"
	"os"
	"os/exec"

	"github.com/rlapz/kvrt_bot_extern/model"
)

func CallApi(api *model.ApiArgs, apiType string, req *model.ApiReq) error {
	text, err := json.Marshal(req)
	if err != nil {
		return err
	}

	cmd := exec.Command(api.Api, api.ConfigFile, api.CmdName, apiType, string(text))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func SendTextPlain(api *model.ApiArgs, text string) error {
	req := model.ApiReq{
		Type:      "plain",
		ChatId:    api.ChatId,
		UserId:    api.UserId,
		MessageId: api.MessageId,
		Deletable: true,
		Text:      text,
	}

	return CallApi(api, "send_text", &req)
}

func SendTextFormat(api *model.ApiArgs, text string) error {
	req := model.ApiReq{
		Type:      "format",
		ChatId:    api.ChatId,
		UserId:    api.UserId,
		MessageId: api.MessageId,
		Deletable: true,
		Text:      text,
	}

	return CallApi(api, "send_text", &req)
}
