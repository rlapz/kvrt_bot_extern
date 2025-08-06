package extra

import (
	"encoding/json"
	"fmt"

	"github.com/rlapz/kvrt_bot_extern/api"
	"github.com/rlapz/kvrt_bot_extern/model"
	"github.com/rlapz/kvrt_bot_extern/util"
)

type stoicism struct {
	Data struct {
		Author string `json:"author"`
		Quote  string `json:"quote"`
	} `json:"data"`
}

func (s *stoicism) fetch() error {
	url := "https://stoic.tekloon.net/stoic-quote"
	body, err := util.FetchGet(url)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, s)
	if err != nil {
		return err
	}

	return nil
}

func (s *stoicism) buildContent() string {
	return fmt.Sprintf("\"_%s_\"\n\n──%s", util.TgEscape(s.Data.Quote), util.TgEscape(s.Data.Author))
}

func RunStoicism(a *model.ApiArgs) {
	var stc stoicism
	err := stc.fetch()
	if err != nil {
		fmt.Println("error:", err)
		_ = api.SendTextPlain(a, err.Error())
		return
	}

	if err = api.SendTextFormat(a, stc.buildContent()); err != nil {
		fmt.Println("error:", err)
		_ = api.SendTextPlain(a, err.Error())
	}
}
