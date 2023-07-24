package minimax

import (
	"context"
	"errors"
	"net/http"
)

// Chat message role defined by the OpenAI API.
const (
	SenderTypeUser = "USER"
	SenderTypeBot  = "BOT"
)

type FinishReason string

const (
	FinishReasonStop          FinishReason = "stop"
	FinishReasonLength        FinishReason = "length"
	FinishReasonMaxOutput     FinishReason = "max_output"
	FinishReasonFunctionCall  FinishReason = "function_call"
	FinishReasonContentFilter FinishReason = "content_filter"
	FinishReasonNull          FinishReason = "null"
)

const chatCompletionsSuffix = "/text/chatcompletion"

var (
	ErrChatCompletionStreamNotSupported = errors.New("streaming is not supported with this method, please use CreateChatCompletionStream") //nolint:lll
)

type ChatCompletionRequest struct {
	Model            string    `json:"model"`
	Prompt           string    `json:"prompt"`
	RoleMeta         RoleMeta  `json:"role_meta,omitempty"`
	Stream           bool      `json:"stream,omitempty"`
	SSE              bool      `json:"use_standard_sse,omitempty"`
	Messages         []Message `json:"messages"`
	Temperature      float64   `json:"temperature,omitempty"`
	TokensToGenerate int       `json:"tokens_to_generate"`
	TopP             float64   `json:"top_p,omitempty"`
	SkipInfoMask     bool      `json:"skip_info_mask,omitempty"`
}

type RoleMeta struct {
	UserName string `json:"user_name"`
	BotName  string `json:"bot_name"`
}

type Message struct {
	SenderType string `json:"sender_type"`
	Text       string `json:"text"`
}

type ChatCompletionResponse struct {
	ID              string                 `json:"id"`
	Created         int64                  `json:"created"`
	Model           string                 `json:"model"`
	Reply           string                 `json:"reply"`
	Choices         []Choice               `json:"choices"`
	InputSensitive  bool                   `json:"input_sensitive,omitempty"`
	OutputSensitive bool                   `json:"output_sensitive"`
	BaseResp        ChatCompletionBaseResp `json:"base_resp,omitempty"`
	Usage           Usage                  `json:"usage,omitempty"`
}

type Choice struct {
	Delta        string       `json:"delta"`
	Index        int          `json:"index,omitempty"`
	FinishReason FinishReason `json:"finish_reason,omitempty"`
}

type ChoiceDelta struct {
	Choices []Choice `json:"choices"`
}

type ChatCompletionBaseResp struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type Usage struct {
	TotalTokens int `json:"total_tokens"`
}

func (c *Client) CreateChatCompletion(
	ctx context.Context,
	request ChatCompletionRequest,
) (response ChatCompletionResponse, err error) {
	if request.Stream {
		err = ErrChatCompletionStreamNotSupported
		return
	}

	urlSuffix := chatCompletionsSuffix

	req, err := c.newRequest(ctx, http.MethodPost, c.fullUrl(urlSuffix, request.Model), withBody(request))
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}
