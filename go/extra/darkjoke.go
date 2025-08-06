package extra

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/rlapz/kvrt_bot_extern/api"
	"github.com/rlapz/kvrt_bot_extern/model"
	"github.com/rlapz/kvrt_bot_extern/util"
)

type darkjoke struct {
	Id       uint64 `json:"id"`
	Type     string `json:"type"`
	Setup    string `json:"setup"`
	Delivery string `json:"delivery"`
	Joke     string `json:"joke"`
	Error    bool   `json:"error"`
}

func (d *darkjoke) fetch(isNsfw bool) error {
	url := "https://v2.jokeapi.dev/joke/Dark?blacklistFlags=religious,racist"
	if !isNsfw {
		url += ",nsfw"
	}

	body, err := util.FetchGet(url)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, d)
	if err != nil {
		return err
	}

	if d.Error {
		return errors.New("failed to fetch")
	}

	return nil
}

func (d *darkjoke) buildContent() string {
	if strings.EqualFold(d.Type, "single") {
		return fmt.Sprintf("%s", util.TgEscape(d.Joke))
	}

	return fmt.Sprintf("\"%s\"\n\nDelivery: ||%s||",
		util.TgEscape(d.Setup),
		util.TgEscape(d.Delivery),
	)
}

func RunDarkJoke(a *model.ApiArgs) {
	isNsfw := false
	if (a.ChatFlags & model.CHAT_FLAG_ALLOW_CMD_NSFW) != 0 {
		isNsfw = true
	}

	var djk darkjoke
	err := djk.fetch(isNsfw)
	if err != nil {
		fmt.Println("error:", err)
		_ = api.SendTextPlain(a, err.Error())
		return
	}

	if err = api.SendTextFormat(a, djk.buildContent()); err != nil {
		fmt.Println("error:", err)
		_ = api.SendTextPlain(a, err.Error())
	}
}
