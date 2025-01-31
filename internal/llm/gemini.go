package llm

import (
	"context"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GeminiClient struct {
	Model *genai.GenerativeModel
}

type ProxyRoundTripper struct {
	APIKey   string
	ProxyURL string
}

func NewGeminiClient(ctx context.Context) *GeminiClient {

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("GEMINI_API_KEY environment variable not set")
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}

	model := client.GenerativeModel("gemini-1.5-flash")

	return &GeminiClient{
		Model: model,
	}
}
