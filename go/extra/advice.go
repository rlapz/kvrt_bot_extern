package extra

import (
	"encoding/json"
	"fmt"

	"github.com/rlapz/kvrt_bot_extern/api"
	"github.com/rlapz/kvrt_bot_extern/model"
	"github.com/rlapz/kvrt_bot_extern/util"
)

type advice struct {
	Slip struct {
		Id     uint64 `json:"id"`
		Advice string `json:"advice"`
	} `json:"slip"`
}

func (a *advice) fetch() error {
	url := "https://api.adviceslip.com/advice"
	body, err := util.FetchGet(url)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, a)
	if err != nil {
		return err
	}

	return nil
}

func (a *advice) buildContent() string {
	return fmt.Sprintf("\"%s\"", util.TgEscape(a.Slip.Advice))
}

func RunAdvice(a *model.ApiArgs) {
	req := model.ApiReq{
		Type:    "acquire",
		ChatId:  a.ChatId,
		UserId:  a.UserId,
		Context: a.CmdName,
	}

	err := api.Submit(a, "session", &req)
	if err != nil {
		_ = api.SendTextPlain(a, "Please wait!")
		return
	}

	defer func() {
		req.Type = "release"
		err := api.Submit(a, "session", &req)
		if err != nil {
			fmt.Println("error:", err)
		}
	}()

	var adv advice
	err = adv.fetch()
	if err != nil {
		fmt.Println("error:", err)
		_ = api.SendTextPlain(a, err.Error())
		return
	}

	if err = api.SendTextFormat(a, adv.buildContent()); err != nil {
		fmt.Println("error:", err)
		_ = api.SendTextPlain(a, err.Error())
	}
}
