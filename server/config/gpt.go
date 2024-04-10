package config

import "github.com/sashabaranov/go-openai"

const (
	OpenAIKey = `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`

	GPTModel = openai.GPT4

	AnalysisPrompt = `create a JSON response like below using the interview transcription between a lawyer and a client that comes after the json object.
	{"questions":[list of legal questions for the lawyer to ask the client after the interview that have not been asked],
	"summary":"summary of the conversation between the client and lawyer",
	"points":"additional points or notes for the lawyer to keep in mind"}`

	AudioFileExtension = `webm`
)
