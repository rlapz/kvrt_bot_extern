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

func (a *aniquote) fetch() error {
	url := "https://api.animechan.io/v1/quotes/random"
	body, err := util.FetchGet(url)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, a)
	if err != nil {
		return err
	}

	if a.Message != "" {
		return errors.New(a.Message)
	}

	if !strings.EqualFold(a.Status, "success") {
		return errors.New("invalid response")
	}

	return nil
}

func (a *aniquote) buildContent() string {
	content := util.TgEscape(a.Data.Content)
	character := util.TgEscape(a.Data.Character.Name)
	anime := util.TgEscape(a.Data.Anime.Name)

	return fmt.Sprintf("\"_%s_\"\n\n──%s \\(%s\\)", content, character, anime)
}

func RunAniquote(a *model.ApiArgs) {
	var anq aniquote
	err := anq.fetch()
	if err != nil {
		fmt.Println("error:", err)
		_ = api.SendTextPlain(a, err.Error())
		return
	}

	if err = api.SendTextFormat(a, anq.buildContent()); err != nil {
		fmt.Println("error:", err)
		_ = api.SendTextPlain(a, err.Error())
	}
}
