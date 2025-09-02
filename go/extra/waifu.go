package extra

import (
	"encoding/json"
	"errors"
	"slices"
	"strings"

	"github.com/rlapz/kvrt_bot_extern/api"
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

func (w *waifu) fetch(filter string, isNsfw bool) error {
	url := "https://api.waifu.pics/sfw/"
	if isNsfw {
		url = "https://api.waifu.pics/nsfw/"
	}

	body, err := util.FetchGet(url + filter)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, w)
	if err != nil {
		return err
	}

	return nil
}

func RunWaifu(a *model.ApiArgs) error {
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

			return api.SendTextFormat(a, text)
		}
	}

	var wf waifu
	err = wf.fetch(filter, isNsfw)
	if err != nil {
		return err
	}

	if strings.HasSuffix(strings.ToLower(wf.Url), ".gif") {
		return api.SendAnimationUrl(a, wf.Url, "")
	}

	return api.SendPhotoUrl(a, wf.Url, "")
}
