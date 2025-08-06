package extra

import (
	"encoding/json"
	"fmt"

	"github.com/rlapz/kvrt_bot_extern/api"
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

func (q *quote) fetch() error {
	url := "https://www.quoterism.com/api/quotes/random"
	body, err := util.FetchGet(url)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, q)
	if err != nil {
		return err
	}

	return nil
}

func (q *quote) buildContent() string {
	return fmt.Sprintf("\"_%s_\"\n\n──%s", util.TgEscape(q.Text), util.TgEscape(q.Author.Name))
}

func RunQuote(a *model.ApiArgs) {
	var qt quote
	err := qt.fetch()
	if err != nil {
		fmt.Println("error:", err)
		_ = api.SendTextPlain(a, err.Error())
		return
	}

	if err = api.SendTextFormat(a, qt.buildContent()); err != nil {
		fmt.Println("error:", err)
		_ = api.SendTextPlain(a, err.Error())
	}
}
