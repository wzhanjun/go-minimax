package minimax

import (
	"context"
	"net/http"
)

type ChatCompletionProStreamResponse struct {
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

type ChatCompletionProStream struct {
	*streamReader[ChatCompletionProStreamResponse]
}

func (c *Client) CreateChatCompletionProStream(
	ctx context.Context,
	request ChatCompletionProRequest,
) (stream *ChatCompletionProStream, err error) {
	urlSuffix := chatCompletionsProSuffix

	request.Stream = true
	req, err := c.newRequest(ctx, http.MethodPost, c.fullUrl(urlSuffix), withBody(request))
	if err != nil {
		return nil, err
	}

	resp, err := sendRequestStream[ChatCompletionProStreamResponse](c, req)
	if err != nil {
		return
	}
	stream = &ChatCompletionProStream{
		streamReader: resp,
	}
	return
}
