package minimax

import (
	"context"
	"net/http"
)

type ChatCompletionStreamResponse struct {
	CommonResponse
	ID              string   `json:"id"`
	Created         int64    `json:"created"`
	Model           string   `json:"model"`
	Reply           string   `json:"reply"`
	Choices         []Choice `json:"choices"`
	InputSensitive  bool     `json:"input_sensitive,omitempty"`
	OutputSensitive bool     `json:"output_sensitive"`
	Usage           Usage    `json:"usage,omitempty"`
}

// ChatCompletionStream
// Note: Perhaps it is more elegant to abstract Stream using generics.
type ChatCompletionStream struct {
	*streamReader[ChatCompletionStreamResponse]
}

// CreateChatCompletionStream — API call to create a chat completion w/ streaming
// support. It sets whether to stream back partial progress. If set, tokens will be
// sent as data-only server-sent events as they become available, with the
// stream terminated by a data: [DONE] message.
func (c *Client) CreateChatCompletionStream(
	ctx context.Context,
	request ChatCompletionRequest,
) (stream *ChatCompletionStream, err error) {
	urlSuffix := chatCompletionsSuffix

	request.Stream = true
	req, err := c.newRequest(ctx, http.MethodPost, c.fullUrl(urlSuffix), withBody(request))
	if err != nil {
		return nil, err
	}

	resp, err := sendRequestStream[ChatCompletionStreamResponse](c, req)
	if err != nil {
		return
	}
	stream = &ChatCompletionStream{
		streamReader: resp,
	}
	return
}
