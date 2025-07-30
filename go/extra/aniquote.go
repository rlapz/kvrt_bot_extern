package extra

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

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
}

func fetchAniquote() (*aniquote, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.animechan.io/v1/quotes/random", http.NoBody)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data aniquote
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
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
		_ = util.SendTextPlain(a, err.Error())
		return
	}

	if err = util.SendTextFormat(a, buildContentAniquote(res)); err != nil {
		fmt.Println("error:", err)
		_ = util.SendTextPlain(a, err.Error())
	}
}
