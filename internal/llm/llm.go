package llm

import (
	"context"

	"github.com/google/generative-ai-go/genai"
)

type LlmService struct {
	client *GeminiClient
}

func NewLlmService(gemini *GeminiClient) *LlmService {
	return &LlmService{
		client: gemini,
	}
}

func (l *LlmService) GetGenerateResponse(ctx context.Context, req string) ([]string, error) {
	resp, err := l.client.Model.GenerateContent(ctx, genai.Text(req))
	if err != nil {
		return []string{}, err
	}

	response := getText(resp)

	return splitMessage(response, 1800), nil
}

func getText(resp *genai.GenerateContentResponse) string {
	var result string
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				switch p := part.(type) {
				case genai.Text:
					result += string(p)
				default:
				}
			}
		}
	}
	return result
}

func splitMessage(message string, maxLength int) []string {
	var splitted []string
	for len(message) > maxLength {
		splitted = append(splitted, message[:maxLength])
		message = message[maxLength:]
	}
	splitted = append(splitted, message)
	return splitted
}
