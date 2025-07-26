package extra

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"slices"

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

var filters = []string{
	"random", "catgirl", "foxgirl", "wolf-girl", "animal-ears", "tail", "tail-with-ribbon",
	"tail-from-under-skirt", "cute", "cuteness-is-justice", "blue-archive", "girl", "young-girl",
	"maid", "maid-uniform", "vtuber", "w-sitting", "lying-down", "hands-forming-a-heart",
	"wink", "valentine", "headphones", "thigh-high-socks", "knee-high-socks", "white-tights",
	"black-tights", "heterochromia", "uniform", "sailor-uniform", "hoodie", "ribbon", "white-hair",
	"blue-hair", "long-hair", "blonde", "blue-eyes", "purple-eyes",
}

func fetch(filter string) (*neko, error) {
	if len(filter) == 0 {
		filter = "random"
	}

	if !slices.Contains(filters, filter) {
		return nil, errors.New("invalid filter")
	}

	req, err := http.NewRequest(http.MethodGet, "https://api.nekosia.cat/api/v1/images/"+filter, http.NoBody)
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

func buildContent(n *neko) string {
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
	ret, err := fetch("")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	req := model.ApiReq{
		Type:      "format",
		ChatId:    a.ChatId,
		UserId:    a.UserId,
		MessageId: a.MessageId,
		Deletable: true,
		Text:      buildContent(ret),
	}

	text, err := json.Marshal(&req)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	cmdArg := []string{
		a.Config,
		a.CmdName,
		"send_text",
		string(text),
	}

	cmd := exec.Command(a.Api, cmdArg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err = cmd.Run(); err != nil {
		fmt.Println("error:", err)
	}
}
