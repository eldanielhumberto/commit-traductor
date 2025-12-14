package translator

import (
	"context"
	"fmt"

	"google.golang.org/genai"
)

const (
	// DefaultModel is the default Gemini model to use for translation
	DefaultModel = "gemini-2.5-flash-lite"

	// PromptTemplate is the template for translation prompts
	PromptTemplate = "Translate the following message to %s language. Only return the translated message, nothing else: %s"
)

// Translator handles message translation using Gemini API
type Translator struct {
	client *genai.Client
	model  string
}

// New creates a new Translator instance
func New(ctx context.Context, apiKey string) (*Translator, error) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create Gemini client: %w", err)
	}

	return &Translator{
		client: client,
		model:  DefaultModel,
	}, nil
}

// Translate translates a message to the target language
func (t *Translator) Translate(ctx context.Context, message, targetLang string) (string, error) {
	if message == "" {
		return "", fmt.Errorf("message cannot be empty")
	}
	if targetLang == "" {
		return "", fmt.Errorf("target language cannot be empty")
	}

	prompt := fmt.Sprintf(PromptTemplate, targetLang, message)

	result, err := t.client.Models.GenerateContent(
		ctx,
		t.model,
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("failed to generate translation: %w", err)
	}

	return result.Text(), nil
}

// Close closes the translator client
func (t *Translator) Close() error {
	// Add cleanup if needed in the future
	return nil
}
