package minimax

import (
	"context"
	"net/http"
)

type T2AStreamResponse struct {
	CommonResponse
	Data struct {
		Audio  string `json:"audio"`
		Status int    `json:"status"`
		Ced    string `json:"ced"`
	} `json:"data"`
	ExtraInfo AudioExtraInfo `json:"extra_info"`
}

type T2AStream struct {
	*streamReader[T2AStreamResponse]
}

func (c *Client) T2AStream(
	ctx context.Context,
	request TextToSpeechRequest,
) (stream *T2AStream, err error) {

	req, err := c.newRequest(ctx, http.MethodPost, c.fullUrl("/tts/stream"), withBody(request))
	if err != nil {
		return nil, err
	}

	resp, err := sendRequestStream[T2AStreamResponse](c, req)
	if err != nil {
		return
	}
	stream = &T2AStream{
		streamReader: resp,
	}
	return
}
