package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"

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

func CallDirectApi(api *model.ApiArgs, tgMethod string, args ...string) error {
	query := ""
	if len(args) > 0 {
		var stb strings.Builder
		stb.WriteString("?")
		for _, v := range args {
			stb.WriteString(v)
			stb.WriteString("&")
		}

		query = strings.TrimSuffix(stb.String(), "&")
	}

	var stb strings.Builder
	stb.WriteString(api.TgApi)
	stb.WriteString("/")
	stb.WriteString(tgMethod)
	stb.WriteString(query)

	fmt.Println(stb.String())

	req, err := http.NewRequest(http.MethodGet, stb.String(), http.NoBody)
	if err != nil {
		return err
	}
	defer req.Body.Close()

	client := &http.Client{}
	_, err = client.Do(req)
	// TODO: handle the response
	return err
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

func SendPhotoUrl(api *model.ApiArgs, photo, text string) error {
	req := model.ApiReq{
		Type:      "url",
		ChatId:    api.ChatId,
		UserId:    api.UserId,
		MessageId: api.MessageId,
		Deletable: true,
		Photo:     photo,
		Text:      text,
	}

	return CallApi(api, "send_photo", &req)
}
