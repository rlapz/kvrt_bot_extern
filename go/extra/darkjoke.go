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

func fetchDarkJoke(isNsfw bool) (*darkjoke, error) {
	url := "https://v2.jokeapi.dev/joke/Dark?blacklistFlags=religious,racist"
	if !isNsfw {
		url += ",nsfw"
	}

	body, err := util.FetchGet(url)
	if err != nil {
		return nil, err
	}

	var data darkjoke
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if data.Error {
		return nil, errors.New("failed to fetch")
	}

	return &data, nil
}

func buildContentDarkJoke(t *darkjoke) string {
	if strings.EqualFold(t.Type, "single") {
		return fmt.Sprintf("%s", util.TgEscape(t.Joke))
	}

	return fmt.Sprintf("\"%s\"\n\nDelivery: ||%s||",
		util.TgEscape(t.Setup),
		util.TgEscape(t.Delivery),
	)
}

func RunDarkJoke(a *model.ApiArgs) {
	isNsfw := false
	if (a.ChatFlags & model.CHAT_FLAG_ALLOW_CMD_NSFW) != 0 {
		isNsfw = true
	}

	ret, err := fetchDarkJoke(isNsfw)
	if err != nil {
		fmt.Println("error:", err)
		_ = api.SendTextPlain(a, err.Error())
		return
	}

	if err = api.SendTextFormat(a, buildContentDarkJoke(ret)); err != nil {
		fmt.Println("error:", err)
		_ = api.SendTextPlain(a, err.Error())
	}
}
