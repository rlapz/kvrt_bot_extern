package extra

import (
	"encoding/json"
	"fmt"

	"github.com/rlapz/kvrt_bot_extern/model"
	"github.com/rlapz/kvrt_bot_extern/util"
)

type tellme struct {
	Id        string `json:"id"`
	Text      string `json:"text"`
	Source    string `json:"source"`
	SourceUrl string `json:"source_url"`
}

func fetchTellMe() (*tellme, error) {
	url := "https://uselessfacts.jsph.pl/api/v2/facts/random"
	body, err := util.FetchGet(url)
	if err != nil {
		return nil, err
	}

	var data tellme
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func buildContentTellme(t *tellme) string {
	return fmt.Sprintf("%s\n\nSource: [%s](%s)",
		util.TgEscape(t.Text),
		util.TgEscape(t.Source),
		t.SourceUrl,
	)
}

func RunTellMe(a *model.ApiArgs) {
	ret, err := fetchTellMe()
	if err != nil {
		fmt.Println("error:", err)
		_ = util.SendTextPlain(a, err.Error())
		return
	}

	if err = util.SendTextFormat(a, buildContentTellme(ret)); err != nil {
		fmt.Println("error:", err)
		_ = util.SendTextPlain(a, err.Error())
	}
}
