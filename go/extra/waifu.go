package extra

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"slices"
	"strconv"
	"strings"

	"github.com/rlapz/kvrt_bot_extern/model"
	"github.com/rlapz/kvrt_bot_extern/util"
)

type waifu struct {
	Url string `json:"url"`
}

var filtersWaifuSfw = []string{
	"waifu", "neko", "shinobu", "megumin", "bully", "cuddle", "cry", "hug", "awoo", "kiss", "lick",
	"pat", "smug", "bonk", "yeet", "blush", "smile", "wave", "highfive", "handhold", "nom", "bite",
	"glomp", "slap", "kill", "kick", "happy", "wink", "poke", "dance", "cringe",
}

var filtersWaifuNsfw = []string{
	"waifu", "neko", "trap", "blowjob",
}

func getFiltersWaifu(isNsfw bool) []string {
	if isNsfw {
		return filtersWaifuNsfw
	}

	return filtersWaifuSfw
}

func validateFilterWaifu(filter string, isNsfw bool) (string, error) {
	if len(filter) == 0 {
		return "waifu", nil
	}

	if !slices.Contains(getFiltersWaifu(isNsfw), filter) {
		return "", errors.New("invalid filter")
	}

	return filter, nil
}

func fetchWaifu(filter string, isNsfw bool) (string, error) {
	url := "https://api.waifu.pics/sfw/"
	if isNsfw {
		url = "https://api.waifu.pics/nsfw/"
	}

	req, err := http.NewRequest(http.MethodGet, url+filter, http.NoBody)
	if err != nil {
		return "", err
	}
	defer req.Body.Close()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data waifu
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	return data.Url, nil
}

func RunWaifu(a *model.ApiArgs) {
	var err error
	filter := "waifu"
	spl := strings.SplitN(a.Text, " ", 2)

	isNsfw := false
	if (a.ChatFlags & model.CHAT_FLAG_ALLOW_CMD_NSFW) != 0 {
		isNsfw = true
	}

	if len(spl) > 1 {
		filter, err = validateFilterWaifu(strings.ToLower(spl[1]), isNsfw)
		if err != nil {
			var bb strings.Builder
			bb.WriteString("Invalid argument\\!\nAvailable arguments:\n`")

			filters := getFiltersWaifu(isNsfw)
			for _, v := range filters {
				bb.WriteString(v)
				bb.WriteString(", ")
			}

			text := bb.String()
			text = strings.TrimSuffix(text, ", ")
			text += "`"
			if err = util.SendTextFormat(a, text); err != nil {
				fmt.Println("error:", err)
			}

			return
		}
	}

	ret, err := fetchWaifu(filter, isNsfw)
	if err != nil {
		fmt.Println("error:", err)
		_ = util.SendTextPlain(a, err.Error())
		return
	}

	if strings.HasSuffix(strings.ToLower(ret), ".gif") {
		args := []string{
			"animation=" + url.QueryEscape(ret),
			"chat_id=" + strconv.FormatInt(a.ChatId, 10),
			"reply_to_message_id=" + strconv.FormatInt(a.MessageId, 10),
		}
		err = util.CallDirectApi(a, "sendAnimation", args...)
	} else {
		err = util.SendPhotoUrl(a, ret, "")
	}

	if err != nil {
		fmt.Println("error:", err)
		_ = util.SendTextPlain(a, err.Error())
	}
}
