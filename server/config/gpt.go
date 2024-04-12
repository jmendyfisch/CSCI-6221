package config

import "github.com/sashabaranov/go-openai"

const (
	OpenAIKey = `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`

	GPTModel = openai.GPT4

	AnalysisPrompt = `create a JSON response like below using the interview transcription between a lawyer and a client keeping in mind the city and state.
	{"questions":[list of questions for the lawyer to ask the client after the interview that have not been asked. these should be questions that will elicit more useful details about the case from the client. avoid asking the client questions about the state of the law. these should just be things the client would know that would be helpful to the lawyer.],
	"summary":"summary of the conversation between the client and lawyer",
	"points":[list of additional points or notes for the lawyer to keep in mind]}`

	CaseSummarizationPrompt = `create a summary of the given legal case using the case description and meeting summaries provided.`
)

var (
	AudioFileExtensions = []string{"webm", "flac", "mp3", "mp4", "mpeg", "mpga", "m4a", "ogg", "wav"}
)
