package minimax

import (
	"context"
	"errors"
	"net/http"
)

type SenderType string

const (
	SenderTypeBot      SenderType = "BOT"
	SenderTypeUser     SenderType = "USER"
	SenderTypeFunction SenderType = "FUNCTION"
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

type SensitiveType int64

const (
	SensitiveTypeSerious     = 1 // 严重违规
	SensitiveTypePron        = 2 // 色情
	SensitiveTypeAdvertising = 3 // 广告
	SensitiveTypeProhibited  = 4 // 违禁
	SensitiveTypeAbuse       = 5 // 谩骂
	SensitiveTypeTerrorism   = 6 // 暴恐
	SensitiveTypeOthers      = 7 // 其他
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
	Temperature      float32   `json:"temperature,omitempty"`
	TokensToGenerate int       `json:"tokens_to_generate"`
	TopP             float32   `json:"top_p,omitempty"`
	SkipInfoMask     bool      `json:"skip_info_mask,omitempty"`
}

type RoleMeta struct {
	UserName string `json:"user_name"`
	BotName  string `json:"bot_name"`
}

type Message struct {
	SenderType SenderType `json:"sender_type"`
	SenderName string     `json:"sender_name,omitempty"`
	Text       string     `json:"text"`
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

type ChatCompletionBaseResp struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type Usage struct {
	TotalTokens           int `json:"total_tokens"`
	TokensWithAddedPlugin int `json:"tokens_with_added_plugin,omitempty"`
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

	req, err := c.newRequest(ctx, http.MethodPost, c.fullUrl(urlSuffix), withBody(request))
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}
