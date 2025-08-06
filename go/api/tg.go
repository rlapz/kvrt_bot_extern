package api

import (
	"fmt"
	"strings"

	"github.com/rlapz/kvrt_bot_extern/model"
	"github.com/rlapz/kvrt_bot_extern/util"
)

func SubmitDirect(api *model.ApiArgs, tgMethod string, args ...string) error {
	query := ""
	if len(args) > 0 {
		var stb strings.Builder
		stb.WriteString("?")
		for _, v := range args {
			stb.WriteString(v)
			stb.WriteString("&")
		}

		query = strings.TrimSuffix(stb.String(), "&")
	}

	url := fmt.Sprintf("%s/%s%s", api.TgApi, tgMethod, query)
	_, err := util.FetchGet(url)

	// TODO: handle the response
	return err
}
