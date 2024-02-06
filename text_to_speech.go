package minimax

import (
	"context"
	"errors"
	"io"
	"net/http"
)

var voiceIdMap = map[string]bool{
	"male-qn-qingse":             true,
	"male-qn-jingying":           true,
	"male-qn-badao":              true,
	"male-qn-daxuesheng":         true,
	"female-shaonv":              true,
	"female-yujie":               true,
	"female-chengshu":            true,
	"female-tianmei":             true,
	"presenter_male":             true,
	"presenter_female":           true,
	"audiobook_male_1":           true,
	"audiobook_male_2":           true,
	"audiobook_female_1":         true,
	"audiobook_female_2":         true,
	"male-qn-qingse-jingpin":     true,
	"male-qn-jingying-jingpin":   true,
	"male-qn-badao-jingpin":      true,
	"male-qn-daxuesheng-jingpin": true,
	"female-shaonv-jingpin":      true,
	"female-yujie-jingpin":       true,
	"female-chengshu-jingpin":    true,
	"female-tianmei-jingpin":     true,
}

type SpeechModel string

const (
	SpeechModel1 SpeechModel = "speech-01"
	SpeechModel2 SpeechModel = "speech-02"
)

type SpeechResponseFormat string

const (
	SpeechResponseFormatMp3  SpeechResponseFormat = "mp3"
	SpeechResponseFormatWav  SpeechResponseFormat = "wav"
	SpeechResponseFormatAac  SpeechResponseFormat = "aac"
	SpeechResponseFormatFlac SpeechResponseFormat = "flac"
	SpeechResponseFormatPcm  SpeechResponseFormat = "pcm"
)

var formatMap = map[SpeechResponseFormat]bool{
	SpeechResponseFormatMp3:  true,
	SpeechResponseFormatWav:  true,
	SpeechResponseFormatAac:  true,
	SpeechResponseFormatFlac: true,
	SpeechResponseFormatPcm:  true,
}

const textToSpeechSuffix = "/text_to_speech"
const t2aproSuffix = "/t2a_pro"

var (
	ErrInvalidSpeechModel = errors.New("invalid speech model")
	ErrInvalidVoice       = errors.New("invalid voice")
)

type TextToSpeechRequest struct {
	Model        string  `json:"model"`
	VoiceId      string  `json:"voice_id"`
	Text         string  `json:"text"`
	Speed        float64 `json:"speed,omitempty"`
	Vol          float64 `json:"vol,omitempty"`
	OutputFormat string  `json:"output_format,omitempty"`
	Pitch        int     `json:"pitch,omitempty"`
}

type T2AProResponse struct {
	CommonResponse
	AudioFile    string         `json:"audio_file"`
	SubtitleFile string         `json:"subtitle_file"`
	ExtraInfo    AudioExtraInfo `json:"extra_info,omitempty"`
}

type AudioExtraInfo struct {
	AudioLength             int64   `json:"audio_length"`
	AudioSampleRate         int64   `json:"audio_sample_rate"`
	AudioSize               int64   `json:"audio_size"`
	BitRate                 int64   `json:"bit_rate"`
	WordCount               int64   `json:"word_count"`
	InvisibleCharacterRatio float64 `json:"invisible_character_ratio"`
}

func contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func isValidSpeechModel(model SpeechModel) bool {
	return contains([]SpeechModel{SpeechModel1, SpeechModel2}, model)
}

func (c *Client) TextToSpeech(ctx context.Context, request TextToSpeechRequest) (content io.ReadCloser, err error) {

	if !isValidSpeechModel(SpeechModel(request.Model)) {
		err = ErrInvalidSpeechModel
		return
	}

	if _, ok := voiceIdMap[request.VoiceId]; !ok {
		err = ErrInvalidVoice
		return
	}
	//
	if _, ok := formatMap[SpeechResponseFormat(request.OutputFormat)]; !ok {
		request.OutputFormat = string(SpeechResponseFormatMp3)
	}

	req, err := c.newRequest(ctx, http.MethodPost, c.fullUrl(textToSpeechSuffix), withBody(request))
	if err != nil {
		return
	}
	req.Header.Set("Accept", "audio/mpeg")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	content, err = c.sendRequestRaw(req)

	return
}

func (c *Client) T2APro(ctx context.Context, request TextToSpeechRequest) (response T2AProResponse, err error) {

	if !isValidSpeechModel(SpeechModel(request.Model)) {
		err = ErrInvalidSpeechModel
		return
	}

	if _, ok := voiceIdMap[request.VoiceId]; !ok {
		err = ErrInvalidVoice
		return
	}
	//
	if _, ok := formatMap[SpeechResponseFormat(request.OutputFormat)]; !ok {
		request.OutputFormat = string(SpeechResponseFormatMp3)
	}

	req, err := c.newRequest(ctx, http.MethodPost, c.fullUrl(t2aproSuffix), withBody(request))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	err = c.sendRequest(req, &response)

	return
}
