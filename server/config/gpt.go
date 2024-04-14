// Contains the OpenAI configurations
// and the Prompts for Transcription Analysis and Case summarization.

package config

import "github.com/sashabaranov/go-openai"

const (
	OpenAIKey = `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`

	GPTModel = openai.GPT3Dot5Turbo0125

	AnalysisPrompt = `create a JSON response like below using the provided transcription, which captures part of a meeting between a lawyer and a client.  The client's city and state, a description of the case written by the client, the lawyer's notes on the present meeting (if any), and summaries of previous conversations between the lawyer and the client (if any exist) are provided. Consider these as background. Make sure it is a properly-formatted JSON response and that there are no invalid characters. (For example, there should be no newline characters.)
	{"questions":[list of questions for the lawyer to ask the client after the interview that have not been asked. these should be questions that will elicit more useful details about the case from the client. avoid asking the client questions about the state of the law. these should just be things the client would know that would be helpful to the lawyer. Format as a comma-separated list of strings, each string is in quotes.],
	"summary":"summary of the conversation between the client and lawyer. focus on the part of the interview contained within the transcription. The summary should capture the most important details of the present conversation. The previous summaries, the client's city and state, and the case description should be considered as background when creating the summary, but there is no need to repeat these details. Don't begin with 'in the present interview,' in the 'recent coversation,' or like filler words -- jump right into the summary. maximum length 150 words.",
	"points":[list of additional points for the lawyer to keep in mind, including possibly relevant laws or areas for further research. End sentences with a period. Format as a comma-separated list of strings, each string is in quotes.]}`

	CaseSummarizationPrompt = `create a summary of the given legal case using the case description and meeting summaries provided. don't start with the word 'summary.' maximum length 200 words.`
)

var (
	AudioFileExtensions = []string{"webm", "flac", "mp3", "mp4", "mpeg", "mpga", "m4a", "ogg", "wav"}
)
