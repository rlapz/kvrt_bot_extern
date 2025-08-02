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

type darkjoke struct {
	Id       uint64 `json:"id"`
	Type     string `json:"type"`
	Setup    string `json:"setup"`
	Delivery string `json:"delivery"`
	Joke     string `json:"joke"`
	Error    bool   `json:"error"`
}

func fetchDarkJoke() (*darkjoke, error) {
	url := "https://v2.jokeapi.dev/joke/Dark?blacklistFlags=nsfw,religious,racist"
	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data darkjoke
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if data.Error {
		return nil, errors.New("failed to fetch")
	}

	return &data, nil
}

func buildContentDarkJoke(t *darkjoke) string {
	if strings.ToLower(t.Type) == "single" {
		return fmt.Sprintf("%s", util.TgEscape(t.Joke))
	}

	return fmt.Sprintf("\"%s\"\n\nDelivery: ||%s||",
		util.TgEscape(t.Setup),
		util.TgEscape(t.Delivery),
	)
}

func RunDarkJoke(a *model.ApiArgs) {
	ret, err := fetchDarkJoke()
	if err != nil {
		fmt.Println("error:", err)
		_ = util.SendTextPlain(a, err.Error())
		return
	}

	if err = util.SendTextFormat(a, buildContentDarkJoke(ret)); err != nil {
		fmt.Println("error:", err)
		_ = util.SendTextPlain(a, err.Error())
	}
}
