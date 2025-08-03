package extra

import (
	"encoding/json"
	"fmt"

	"github.com/rlapz/kvrt_bot_extern/model"
	"github.com/rlapz/kvrt_bot_extern/util"
)

type stoicism struct {
	Data struct {
		Author string `json:"author"`
		Quote  string `json:"quote"`
	} `json:"data"`
}

func fetchStoicism() (*stoicism, error) {
	url := "https://stoic.tekloon.net/stoic-quote"
	body, err := util.FetchGet(url)
	if err != nil {
		return nil, err
	}

	var data stoicism
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func buildContentStoicism(t *stoicism) string {
	return fmt.Sprintf("\"_%s_\"\n\n──%s", util.TgEscape(t.Data.Quote), util.TgEscape(t.Data.Author))
}

func RunStoicism(a *model.ApiArgs) {
	ret, err := fetchStoicism()
	if err != nil {
		fmt.Println("error:", err)
		_ = util.SendTextPlain(a, err.Error())
		return
	}

	if err = util.SendTextFormat(a, buildContentStoicism(ret)); err != nil {
		fmt.Println("error:", err)
		_ = util.SendTextPlain(a, err.Error())
	}
}
