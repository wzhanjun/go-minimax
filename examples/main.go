package main

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/wzhanjun/go-minimax"
)

var token = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJOYW1lIjoidGVzdCIsIlN1YmplY3RJRCI6IjE2ODk2NDk0MTU1NjQ0MjEiLCJQaG9uZSI6Ik1UTXdOemM0TURJME5URT0iLCJHcm91cElEIjoiMTY4OTY0OTQxNTA3NDAzMSIsIlBhZ2VOYW1lIjoiIiwiTWFpbCI6ImFhYXNheW9rQGdtYWlsLmNvbSIsIkNyZWF0ZVRpbWUiOiIyMDIzLTA3LTE4IDE2OjQyOjMwIiwiaXNzIjoibWluaW1heCJ9.dNy8oFhrN8QI1Nxjr5Ug9RZhUCH8ak_jxnZVYx1GxnRm45IGIga-hsrxm2FMnw-oAHe1g3j4Nk5VFIyYKSBqZ42tdLBFkxiMHtfLNi7yV2tZWUI819wtQS5ByW32WjUd8PscQ_V0VexL78Jy5a-sOMu0Kzv2UOrRpqF89toWEtg3GgCZ6eMLCNgOepbBniLw-iQ7hEbnxK6GmtjQUGZLvQLsZORPdKYi94KFJSFay2t3dYhXN0UkICgnR1qUsQ8wb18FQtYRIbIkFsq59eS-pCIdSk9SYViT9Z7Q7EpKfr8T17F_6-0DVSDNSAZN8KnHp-8jZTl4cDU7G124i1haNA"
var orgId = "1689649415074031"
var client = minimax.NewClient(token, orgId)

func main() {
	// chatcompletion()
	chatcompletionstream()
}

func chatcompletion() {
	resp, err := client.CreateChatCompletion(context.Background(), minimax.ChatCompletionRequest{
		Model:            minimax.MiniMaxChat55,
		Prompt:           "你是一个诗人",
		TokensToGenerate: 512,
		Messages: []minimax.Message{
			{
				SenderType: minimax.SenderTypeUser,
				Text:       "模仿李白写一首诗",
			},
		},
		RoleMeta: minimax.RoleMeta{
			UserName: "user",
			BotName:  "ai",
		},
	})

	fmt.Printf("%+v, %+v \n", resp, err)
}

func chatcompletionstream() {
	stream, err := client.CreateChatCompletionStream(context.Background(), minimax.ChatCompletionRequest{
		Model:            minimax.MiniMaxChat55,
		Prompt:           "你是一个诗人",
		TokensToGenerate: 512,
		Messages: []minimax.Message{
			{
				SenderType: minimax.SenderTypeUser,
				Text:       "模仿李白写一首诗",
			},
		},
		RoleMeta: minimax.RoleMeta{
			UserName: "user",
			BotName:  "ai",
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		resp, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		fmt.Printf("%#v\n", resp)
	}
}
