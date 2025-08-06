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

type aniquote struct {
	Status string `json:"status"`
	Data   struct {
		Content string `json:"content"`
		Anime   struct {
			Id   uint64 `json:"id"`
			Name string `json:"name"`
		} `json:"anime"`
		Character struct {
			Id   uint64 `json:"id"`
			Name string `json:"name"`
		} `json:"character"`
	} `json:"data"`
	Message string `json:"message"`
}

func fetchAniquote() (*aniquote, error) {
	url := "https://api.animechan.io/v1/quotes/random"
	body, err := util.FetchGet(url)
	if err != nil {
		return nil, err
	}

	var data aniquote
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if data.Message != "" {
		return nil, errors.New(data.Message)
	}

	if !strings.EqualFold(data.Status, "success") {
		return nil, errors.New("invalid response")
	}

	return &data, nil
}

func buildContentAniquote(a *aniquote) string {
	content := util.TgEscape(a.Data.Content)
	character := util.TgEscape(a.Data.Character.Name)
	anime := util.TgEscape(a.Data.Anime.Name)

	return fmt.Sprintf("\"_%s_\"\n\n──%s \\(%s\\)", content, character, anime)
}

func RunAniquote(a *model.ApiArgs) {
	res, err := fetchAniquote()
	if err != nil {
		fmt.Println("error:", err)
		_ = api.SendTextPlain(a, err.Error())
		return
	}

	if err = api.SendTextFormat(a, buildContentAniquote(res)); err != nil {
		fmt.Println("error:", err)
		_ = api.SendTextPlain(a, err.Error())
	}
}
