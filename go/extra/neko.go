package extra

import (
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/rlapz/kvrt_bot_extern/model"
	"github.com/rlapz/kvrt_bot_extern/util"
)

type neko struct {
	Success bool `json:"success"`

	Image struct {
		Original struct {
			Url string `json:"url"`
		} `json:"original"`
		Compressed struct {
			Url string `json:"url"`
		} `json:"compressed"`
	} `json:"image"`

	Category string `json:"category"`

	Anime struct {
		Title     string `json:"title"`
		Character string `json:"character"`
	} `json:"anime"`

	Source struct {
		Url string `json:"url"`
	} `json:"source"`

	Attribution struct {
		Artist struct {
			Username string `json:"username"`
			Profile  string `json:"profile"`
		} `jsono:"artist"`
	} `json:"attribution"`
}

var filtersNeko = []string{
	"random", "catgirl", "foxgirl", "wolf-girl", "animal-ears", "tail", "tail-with-ribbon",
	"tail-from-under-skirt", "cute", "cuteness-is-justice", "blue-archive", "girl", "young-girl",
	"maid", "maid-uniform", "vtuber", "w-sitting", "lying-down", "hands-forming-a-heart",
	"wink", "valentine", "headphones", "thigh-high-socks", "knee-high-socks", "white-tights",
	"black-tights", "heterochromia", "uniform", "sailor-uniform", "hoodie", "ribbon", "white-hair",
	"blue-hair", "long-hair", "blonde", "blue-eyes", "purple-eyes",
}

func validateFilterNeko(filter string) (string, error) {
	if len(filter) == 0 {
		return "random", nil
	}

	if !slices.Contains(filtersNeko, filter) {
		return "", errors.New("invalid filter")
	}

	return filter, nil
}

func fetchNeko(filter string) (*neko, error) {
	url := "https://api.nekosia.cat/api/v1/images/" + filter
	body, err := util.FetchGet(url)
	if err != nil {
		return nil, err
	}

	var data neko
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if !data.Success {
		return nil, errors.New("invalid response")
	}

	return &data, nil
}

func buildContentNeko(n *neko) string {
	anime_chr := n.Anime.Character
	if len(anime_chr) == 0 {
		anime_chr = "?"
	}

	anime_title := n.Anime.Title
	if len(anime_title) == 0 {
		anime_title = "?"
	}

	artist_uname := n.Attribution.Artist.Username
	if len(artist_uname) == 0 {
		artist_uname = "?"
	}

	artist_profile := n.Attribution.Artist.Profile
	if len(artist_profile) == 0 {
		artist_profile = "?"
	}

	return fmt.Sprintf("`URL     :` [Compressed](%s) \\- [Original](%s)\n"+
		"`Name    : %s` from `%s`\n"+
		"`Artist  :`  [%s](%s)\n"+
		"`Source  :`  %s\n"+
		"`Category: %s`",
		n.Image.Compressed.Url,
		n.Image.Original.Url,
		anime_chr,
		anime_title,
		util.TgEscape(artist_uname),
		artist_profile,
		util.TgEscape(n.Source.Url),
		n.Category,
	)
}

func RunNeko(a *model.ApiArgs) {
	var err error
	filter := "random"
	spl := strings.SplitN(a.Text, " ", 2)
	if len(spl) > 1 {
		filter, err = validateFilterNeko(strings.ToLower(spl[1]))
		if err != nil {
			var bb strings.Builder
			bb.WriteString("Invalid argument\\!\nAvailable arguments:\n`")
			for _, v := range filtersNeko {
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

	ret, err := fetchNeko(filter)
	if err != nil {
		fmt.Println("error:", err)
		_ = util.SendTextPlain(a, err.Error())
		return
	}

	if err = util.SendPhotoUrl(a, ret.Image.Compressed.Url, buildContentNeko(ret)); err != nil {
		fmt.Println("error:", err)
		_ = util.SendTextPlain(a, err.Error())
	}
}
