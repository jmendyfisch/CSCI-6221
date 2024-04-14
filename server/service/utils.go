package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"server/config"
	"server/database"
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

func getOutputTextFromTranscription(caseID int, meetingID int, text string, gptSums []string) (res types.GPTPromptOutput, err error) {

	var finalText string

	MeetingDetails, _ := database.GetMeetingDetails(meetingID)

	desc, city, state, err := getDetails(caseID)
	if err != nil {
		return
	}

	finalText = "Case description provided by client: " + desc + "\nClient city and state: " + city + ", " + state + "\n Lawyer notes from meeting: " + MeetingDetails.LawyerNotes.String + "\n Summaries of previous conversations between the lawyer and client: "

	for _, sum := range gptSums {
		finalText += "\n" + sum
	}

	finalText = finalText + "\n*** The interview transcription for the present conversation starts here: " + text

	fmt.Println("INPUT TEXT FOR CASE SUMMARY: ", finalText)

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
				Content: finalText,
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

func getCaseSummary(desc string, gptSums []string) (string, error) {
	text := "Client's description of case - " + desc + "\n meeting summaries - "

	for _, sum := range gptSums {
		text += "\n" + sum
	}

	fmt.Println("INPUT TEXT FOR CASE SUMMARY: ", text)

	client := openai.NewClient(config.OpenAIKey)

	req := openai.ChatCompletionRequest{
		Model: config.GPTModel,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: config.CaseSummarizationPrompt,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: text,
			},
		},
	}

	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		log.Println("ChatGPT Case Summary Completion error: ", err)
		return "", err
	}

	jsonText := resp.Choices[0].Message.Content

	log.Println("case summary json resp from gpt: \n", jsonText)

	return jsonText, nil
}

func getDetails(caseID int) (description, city, state string, err error) {
	c, err := database.GetCaseDetails(caseID)
	if err != nil {
		return
	}

	return c.Description, c.AddressCity, c.AddressState, nil
}
