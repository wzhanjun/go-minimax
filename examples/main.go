package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/wzhanjun/go-minimax"
)

var token = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJOYW1lIjoidGVzdCIsIlN1YmplY3RJRCI6IjE2ODk2NDk0MTU1NjQ0MjEiLCJQaG9uZSI6Ik1UTXdOemM0TURJME5URT0iLCJHcm91cElEIjoiMTY4OTY0OTQxNTA3NDAzMSIsIlBhZ2VOYW1lIjoiIiwiTWFpbCI6ImFhYXNheW9rQGdtYWlsLmNvbSIsIkNyZWF0ZVRpbWUiOiIyMDIzLTA3LTE4IDE2OjQyOjMwIiwiaXNzIjoibWluaW1heCJ9.dNy8oFhrN8QI1Nxjr5Ug9RZhUCH8ak_jxnZVYx1GxnRm45IGIga-hsrxm2FMnw-oAHe1g3j4Nk5VFIyYKSBqZ42tdLBFkxiMHtfLNi7yV2tZWUI819wtQS5ByW32WjUd8PscQ_V0VexL78Jy5a-sOMu0Kzv2UOrRpqF89toWEtg3GgCZ6eMLCNgOepbBniLw-iQ7hEbnxK6GmtjQUGZLvQLsZORPdKYi94KFJSFay2t3dYhXN0UkICgnR1qUsQ8wb18FQtYRIbIkFsq59eS-pCIdSk9SYViT9Z7Q7EpKfr8T17F_6-0DVSDNSAZN8KnHp-8jZTl4cDU7G124i1haNA"
var orgId = "1689649415074031"
var client = minimax.NewClient(token, orgId)

func main() {
	// chatcompletion()
	// chatcompletionstream()
	// chatcompletionpro()
	// chatcompletionprostream()
	textToSpeech()

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

func chatcompletionpro() {
	resp, err := client.CreateChatCompletionPro(context.Background(), minimax.ChatCompletionProRequest{
		Model:            minimax.MiniMaxChat55,
		TokensToGenerate: 512,
		Messages: []minimax.Message{
			{
				SenderType: minimax.SenderTypeUser,
				SenderName: "张三",
				Text:       "今天深圳的天气怎么样",
			},
		},
		BotSetting: []minimax.BotSetting{
			{
				BotName: "AI",
				Content: "智能机器人",
			},
		},
		ReplyConstraints: minimax.ReplyConstraints{
			SenderType: minimax.SenderTypeBot,
			SenderName: "AI",
		},
		Plugins: []minimax.PluginName{minimax.PluginNameWebSearch},
	})

	fmt.Printf("%+v, %+v \n", resp, err)
}

func chatcompletionprostream() {
	stream, err := client.CreateChatCompletionProStream(context.Background(), minimax.ChatCompletionProRequest{
		Model:            minimax.MiniMaxChat55,
		TokensToGenerate: 512,
		Messages: []minimax.Message{
			{
				SenderType: minimax.SenderTypeUser,
				SenderName: "张三",
				Text:       "今天深圳的天气怎么样",
			},
		},
		BotSetting: []minimax.BotSetting{
			{
				BotName: "AI",
				Content: "智能机器人",
			},
		},
		ReplyConstraints: minimax.ReplyConstraints{
			SenderType: minimax.SenderTypeBot,
			SenderName: "AI",
		},
		Plugins: []minimax.PluginName{minimax.PluginNameWebSearch},
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

func textToSpeech() {
	resp, err := client.TextToSpeech(context.Background(), minimax.TextToSpeechRequest{
		Model:   minimax.MiniMaxSpeech01,
		VoiceId: "male-qn-qingse",
		Text:    "在一个宁静的小镇，有一座古老的豪宅，据说每到夜晚，豪宅内都会传来神秘的低语和异样的声音。当地居民都对这座豪宅敬而远之，因为传说这里隐藏着一个可怕的存在。",
	})
	if err != nil {
		fmt.Printf("text to audio failed, err:%+v", err)
		return
	}

	file, err := os.Create("audio.mp3")
	if err != nil {
		fmt.Println("创建文件失败:", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp)
	if err != nil {
		fmt.Println("写入文件失败:", err)
		return
	}

	fmt.Printf("%+v, %+v \n", resp, err)
}
