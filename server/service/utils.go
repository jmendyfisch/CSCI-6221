package service

import (
	"context"
	"encoding/json"
	"log"
	"server/config"
	"server/types"

	"github.com/sashabaranov/go-openai"
)

// Get transcription of the given audio file.
func getTextFromAudio(audioFilePath string) (s string, err error) {

	client := openai.NewClient(config.OpenAIKey)
	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: audioFilePath,
	}

	resp, err := client.CreateTranscription(context.Background(), req)
	if err != nil {
		log.Println("Error getting transcription, err: ", err)
		return "", err
	}

	return resp.Text, nil
}

func getOutputTextFromTranscription(text string) (res types.GPTPromptOutput, err error) {

	client := openai.NewClient(config.OpenAIKey)

	req := openai.ChatCompletionRequest{
		Model: config.GPTModel,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: config.AnalysisPrompt,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: text,
			},
		},
	}

	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		log.Println("ChatGPT Completion error: ", err)
		return
	}

	jsonText := resp.Choices[0].Message.Content

	log.Println("json resp from gpt: \n", jsonText)

	err = json.Unmarshal([]byte(jsonText), &res)
	if err != nil {
		log.Println("ChatGPT JSON response unmarshal error: ", err)
	}

	return res, nil
}
