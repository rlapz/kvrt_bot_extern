package extra

import (
	"encoding/json"
	"fmt"

	"github.com/rlapz/kvrt_bot_extern/api"
	"github.com/rlapz/kvrt_bot_extern/model"
	"github.com/rlapz/kvrt_bot_extern/util"
)

type joke struct {
	Id        uint64 `json:"id"`
	Type      string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

func fetchJoke() (*joke, error) {
	url := "https://official-joke-api.appspot.com/random_joke"
	body, err := util.FetchGet(url)
	if err != nil {
		return nil, err
	}

	var data joke
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func buildContentJoke(t *joke) string {
	return fmt.Sprintf("\"%s\"\n\nAnswer: ||%s||",
		util.TgEscape(t.Setup),
		util.TgEscape(t.Punchline),
	)
}

func RunJoke(a *model.ApiArgs) {
	ret, err := fetchJoke()
	if err != nil {
		fmt.Println("error:", err)
		_ = api.SendTextPlain(a, err.Error())
		return
	}

	if err = api.SendTextFormat(a, buildContentJoke(ret)); err != nil {
		fmt.Println("error:", err)
		_ = api.SendTextPlain(a, err.Error())
	}
}
