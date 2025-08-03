package extra

import (
	"encoding/json"
	"fmt"

	"github.com/rlapz/kvrt_bot_extern/model"
	"github.com/rlapz/kvrt_bot_extern/util"
)

type quote struct {
	Id     string `json:"id"`
	Text   string `json:"text"`
	Author struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"author"`
}

func fetchQuote() (*quote, error) {
	url := "https://www.quoterism.com/api/quotes/random"
	body, err := util.FetchGet(url)
	if err != nil {
		return nil, err
	}

	var data quote
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func buildContentQuote(t *quote) string {
	return fmt.Sprintf("\"_%s_\"\n\n──%s", util.TgEscape(t.Text), util.TgEscape(t.Author.Name))
}

func RunQuote(a *model.ApiArgs) {
	ret, err := fetchQuote()
	if err != nil {
		fmt.Println("error:", err)
		_ = util.SendTextPlain(a, err.Error())
		return
	}

	if err = util.SendTextFormat(a, buildContentQuote(ret)); err != nil {
		fmt.Println("error:", err)
		_ = util.SendTextPlain(a, err.Error())
	}
}
