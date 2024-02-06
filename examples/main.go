package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/wzhanjun/go-minimax"
)

var token = os.Getenv("ENV_MINIMAX_TOKEN")
var orgId = os.Getenv("ENV_MINIMAX_ORG_ID")
var client = minimax.NewClient(token, orgId)

func main() {
	chatcompletion()
	// chatcompletionstream()
	// chatcompletionpro()
	// chatcompletionprostream()
	// textToSpeech()
	// t2APro()

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

func t2APro() {
	resp, err := client.T2APro(context.Background(), minimax.TextToSpeechRequest{
		Model:   minimax.MiniMaxSpeech01,
		VoiceId: "male-qn-qingse",
		Text:    "在一个宁静的小镇，有一座古老的豪宅，据说每到夜晚，豪宅内都会传来神秘的低语和异样的声音。当地居民都对这座豪宅敬而远之，因为传说这里隐藏着一个可怕的存在。",
	})
	if err != nil {
		fmt.Printf("text to audio failed, err:%+v", err)
		return
	}

	fmt.Printf("%+v, %+v \n", resp, err)
}
