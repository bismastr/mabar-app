package llm

import (
	"context"
	"log"
	"net/http"
	"net/url"
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
	c := &http.Client{Transport: &ProxyRoundTripper{
		APIKey: os.Getenv("GEMINI_API_KEY"),
	}}

	client, err := genai.NewClient(ctx, option.WithHTTPClient(c))
	if err != nil {
		log.Fatal(err)
	}

	model := client.GenerativeModel("gemini-1.5-pro")

	return &GeminiClient{
		Model: model,
	}
}

func (t *ProxyRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	transport := http.DefaultTransport.(*http.Transport).Clone()

	if t.ProxyURL != "" {
		proxyURL, err := url.Parse(t.ProxyURL)
		if err != nil {
			return nil, err
		}
		transport.Proxy = http.ProxyURL(proxyURL)
	}

	newReq := req.Clone(req.Context())
	vals := newReq.URL.Query()
	vals.Set("key", t.APIKey)
	newReq.URL.RawQuery = vals.Encode()

	resp, err := transport.RoundTrip(newReq)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
