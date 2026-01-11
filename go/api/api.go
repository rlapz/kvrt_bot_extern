package api

import (
	"encoding/json"
	"os"
	"os/exec"

	"github.com/rlapz/kvrt_bot_extern/model"
)

func Submit(api *model.ApiArgs, apiType string, req *model.ApiReq) error {
	text, err := json.Marshal(req)
	if err != nil {
		return err
	}

	cmd := exec.Command(api.Api, api.ConfigFile, api.CmdName, apiType, string(text))
	cmd.Env = append(cmd.Env, "TG_DB_MAIN_FILE="+api.DbMainFile)
	cmd.Env = append(cmd.Env, "TG_DB_SCHED_FILE="+api.DbSchedFile)
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
		Text:      text,
	}

	return Submit(api, "send_text", &req)
}

func SendTextFormat(api *model.ApiArgs, text string) error {
	req := model.ApiReq{
		Type:      "format",
		ChatId:    api.ChatId,
		UserId:    api.UserId,
		MessageId: api.MessageId,
		Text:      text,
	}

	return Submit(api, "send_text", &req)
}

func SendPhotoUrl(api *model.ApiArgs, photo, text string) error {
	req := model.ApiReq{
		ChatId:    api.ChatId,
		UserId:    api.UserId,
		MessageId: api.MessageId,
		Photo:     photo,
		Text:      text,
		TextType:  "format",
	}

	return Submit(api, "send_photo", &req)
}

func SendAnimationUrl(api *model.ApiArgs, animation, text string) error {
	req := model.ApiReq{
		ChatId:    api.ChatId,
		UserId:    api.UserId,
		MessageId: api.MessageId,
		Animation: animation,
		Text:      text,
		TextType:  "format",
	}

	return Submit(api, "send_animation", &req)
}
