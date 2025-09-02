package extra

import (
	"encoding/json"
	"fmt"

	"github.com/rlapz/kvrt_bot_extern/api"
	"github.com/rlapz/kvrt_bot_extern/model"
	"github.com/rlapz/kvrt_bot_extern/util"
)

type tellme struct {
	Id        string `json:"id"`
	Text      string `json:"text"`
	Source    string `json:"source"`
	SourceUrl string `json:"source_url"`
}

func (t *tellme) fetch() error {
	url := "https://uselessfacts.jsph.pl/api/v2/facts/random"
	body, err := util.FetchGet(url)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, t)
	if err != nil {
		return err
	}

	return nil
}

func (t *tellme) buildContent() string {
	return fmt.Sprintf("%s\n\nSource: [%s](%s)",
		util.TgEscape(t.Text),
		util.TgEscape(t.Source),
		t.SourceUrl,
	)
}

func RunTellMe(a *model.ApiArgs) error {
	var tlm tellme
	err := tlm.fetch()
	if err != nil {
		return err
	}

	return api.SendTextFormat(a, tlm.buildContent())
}
