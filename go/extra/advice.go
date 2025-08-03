package extra

import (
	"encoding/json"
	"fmt"

	"github.com/rlapz/kvrt_bot_extern/model"
	"github.com/rlapz/kvrt_bot_extern/util"
)

type advice struct {
	Slip struct {
		Id     uint64 `json:"id"`
		Advice string `json:"advice"`
	} `json:"slip"`
}

func fetchAdvice() (*advice, error) {
	url := "https://api.adviceslip.com/advice"
	body, err := util.FetchGet(url)
	if err != nil {
		return nil, err
	}

	var data advice
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func buildContentAdvice(t *advice) string {
	return fmt.Sprintf("\"%s\"", util.TgEscape(t.Slip.Advice))
}

func RunAdvice(a *model.ApiArgs) {
	ret, err := fetchAdvice()
	if err != nil {
		fmt.Println("error:", err)
		_ = util.SendTextPlain(a, err.Error())
		return
	}

	if err = util.SendTextFormat(a, buildContentAdvice(ret)); err != nil {
		fmt.Println("error:", err)
		_ = util.SendTextPlain(a, err.Error())
	}
}
