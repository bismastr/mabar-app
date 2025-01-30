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

func (l *LlmService) GetGenerateResponse(ctx context.Context, req string) (*genai.Part, error) {
	resp, err := l.client.Model.GenerateContent(ctx, genai.Text("7 tambah 8 berapa?"))
	if err != nil {
		return nil, err
	}

	return getText(resp), nil
}

func getText(resp *genai.GenerateContentResponse) *genai.Part {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				return &part
			}
		}
	}

	return nil
}
