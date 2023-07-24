package minimax

import (
	"context"
	"errors"
	"io"
	"net/http"
)

var voiceIdMap = map[string]bool{
	"male-qn-qingse":     true,
	"male-qn-jingying":   true,
	"male-qn-badao":      true,
	"male-qn-daxuesheng": true,
	"female-shaonv":      true,
	"female-yujie":       true,
	"female-chengshu":    true,
	"female-tianmei":     true,
	"presenter_male":     true,
	"presenter_female":   true,
	"audiobook_male_1":   true,
	"audiobook_male_2":   true,
	"audiobook_female_1": true,
	"audiobook_female_2": true,
}

const textToSpeechSuffix = "/text_to_speech"

var (
	ErrVoiceIdNotSupported = errors.New("voice_id not supported with this method") //nolint:lll
)

type TextToSpeechRequest struct {
	Model   string `json:"model"`
	VoiceId string `json:"voice_id"`
	Text    string `json:"text"`
}

func (c *Client) TextToSpeech(
	ctx context.Context,
	request TextToSpeechRequest,
) (content io.ReadCloser, err error) {

	if _, ok := voiceIdMap[request.VoiceId]; !ok {
		err = ErrVoiceIdNotSupported
		return
	}
	urlSuffix := textToSpeechSuffix

	req, err := c.newRequest(ctx, http.MethodPost, c.fullUrl(urlSuffix), withBody(request))
	if err != nil {
		return
	}
	req.Header.Set("Accept", "audio/mpeg")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	content, err = c.sendRequestRaw(req)

	return
}
