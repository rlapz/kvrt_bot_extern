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

func (j *joke) fetch() error {
	url := "https://official-joke-api.appspot.com/random_joke"
	body, err := util.FetchGet(url)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, j)
	if err != nil {
		return err
	}

	return nil
}

func (j *joke) buildContent() string {
	return fmt.Sprintf("\"%s\"\n\nAnswer: ||%s||",
		util.TgEscape(j.Setup),
		util.TgEscape(j.Punchline),
	)
}

func RunJoke(a *model.ApiArgs) {
	var jk joke
	err := jk.fetch()
	if err != nil {
		fmt.Println("error:", err)
		_ = api.SendTextPlain(a, err.Error())
		return
	}

	if err = api.SendTextFormat(a, jk.buildContent()); err != nil {
		fmt.Println("error:", err)
		_ = api.SendTextPlain(a, err.Error())
	}
}
