package minimax

import (
	"context"
	"net/http"
)

const chatCompletionsProSuffix = "/text/chatcompletion_pro"

type PluginName string

const (
	PluginNameWebSearch PluginName = "plugin_web_search"
)

type ChatCompletionProRequest struct {
	Model             string           `json:"model"`
	Stream            bool             `json:"stream"`
	TokensToGenerate  int              `json:"tokens_to_generate"`
	Temperature       float32          `json:"temperature,omitempty"`
	TopP              float32          `json:"top_p,omitempty"`
	MaskSensitiveInfo bool             `json:"mask_sensitive_info"`
	Messages          []Message        `json:"messages"`
	BotSetting        []BotSetting     `json:"bot_setting"`
	ReplyConstraints  ReplyConstraints `json:"reply_constraints"`
	Functions         []Function       `json:"functions,omitempty"`
	Plugins           []PluginName     `json:"plugins,omitempty"`
	SampleMessages    []Message        `json:"sample_messages,omitempty"`
}

type BotSetting struct {
	BotName string `json:"bot_name"`
	Content string `json:"content"`
}

type ReplyConstraints struct {
	SenderType SenderType  `json:"sender_type"`
	SenderName string      `json:"sender_name"`
	Glyph      *ReplyGlyph `json:"glyph,omitempty"`
}

type ReplyGlyph struct {
	Type     string `json:"type"`
	RawGlpyh string `json:"raw_glpyh"`
}

type Function struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Parameters  string `json:"parameters"`
}

type ChatCompletionProResponse struct {
	CommonResponse
	ID                  string        `json:"id"`
	Created             int64         `json:"created"`
	Model               string        `json:"model"`
	Reply               string        `json:"reply"`
	InputSensitive      bool          `json:"input_sensitive,omitempty"`
	InputSensitiveType  SensitiveType `json:"input_sensitive_type,omitempty"`
	OutputSensitive     bool          `json:"output_sensitive,omitempty"`
	OutputSensitiveType SensitiveType `json:"output_sensitive_type,omitempty"`
	Choices             []ChoicePro   `json:"choices"`
	Usage               Usage         `json:"usage,omitempty"`
}

type ChoicePro struct {
	Index        int64        `json:"index,omitempty"`
	Messages     []Message    `json:"messages"`
	FinishReason FinishReason `json:"finish_reason,omitempty"`
}

func (c *Client) CreateChatCompletionPro(
	ctx context.Context,
	request ChatCompletionProRequest,
) (response ChatCompletionProResponse, err error) {
	if request.Stream {
		err = ErrChatCompletionStreamNotSupported
		return
	}

	urlSuffix := chatCompletionsProSuffix

	req, err := c.newRequest(ctx, http.MethodPost, c.fullUrl(urlSuffix), withBody(request))
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}
