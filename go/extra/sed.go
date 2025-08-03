package extra

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/rlapz/kvrt_bot_extern/model"
	"github.com/rlapz/kvrt_bot_extern/tg"
	"github.com/rlapz/kvrt_bot_extern/util"
)

type sed struct {
	cmd string
	old string
	new string
}

func (s *sed) parse(src string) error {
	spl := strings.SplitN(src, " ", 3)
	if len(spl) <= 2 {
		return errors.New("invalid argument")
	}

	s.new = strings.TrimSpace(spl[2])
	if s.new == "" {
		return errors.New("invalid argument")
	}

	s.cmd = spl[0]
	s.old = spl[1]
	return nil
}

func (s *sed) replaceAll(text string) string {
	return strings.ReplaceAll(strings.ToLower(text), strings.ToLower(s.old), s.new)
}

func getText(jsn string) (string, error) {
	var tgu tg.Update
	err := json.Unmarshal([]byte(jsn), &tgu)
	if err != nil {
		return "", err
	}

	if tgu.Message.ReplyTo == nil {
		return "", errors.New("invalid")
	}

	text := tgu.Message.ReplyTo.Text
	if text == "" {
		return "", errors.New("no text field")
	}

	return strings.TrimSpace(text), nil
}

func RunSed(a *model.ApiArgs) {
	var s sed
	err := s.parse(a.Text)
	if err != nil {
		fmt.Println("error:", err)
		_ = util.SendTextPlain(a, err.Error())
		return
	}

	orig, err := getText(a.RawJSON)
	if err != nil {
		fmt.Println("error:", err)
		_ = util.SendTextPlain(a, err.Error())
		return
	}

	res := s.replaceAll(orig)
	res = fmt.Sprintf("_Did you mean:_\n\n%s", util.TgEscape(res))

	err = util.SendTextFormat(a, res)
	if err != nil {
		fmt.Println("error:", err)
		_ = util.SendTextPlain(a, err.Error())
	}
}
