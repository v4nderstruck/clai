package gemini

import (
	"context"
	"fmt"
	"os"

	"github.com/v4nderstruck/clai/internal/models"
	"google.golang.org/genai"
)

type GeminiModel struct{}

func NewGeminiModel() *GeminiModel {
	return &GeminiModel{}
}

func (g *GeminiModel) ModelHelp() string {
  if os.Getenv("GEMINI_API_KEY") != "" {
    return "GEMINI_API_KEY set. Using Gemini Model Family."
  }
  return "GEMINI_API_KEY not set. Make sure to set this to use the gemini Model Family."

}

func (g *GeminiModel) OneShotPrompt(thinkLevel models.ThinkingLevel, systemPrompt string, prompt string) (string, error) {
	var err error
	ctx := context.Background()

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return "", fmt.Errorf("cannot get gemini client %v", err)
	}
	config := &genai.GenerateContentConfig{
		SystemInstruction: genai.NewContentFromText(systemPrompt, genai.RoleUser),
	}

	var modelName string

	switch thinkLevel {

	case models.FastResponse:
		modelName = "gemini-2.5-flash-preview-04-17"
	case models.NormalResponse:
		modelName = "gemini-2.5-pro-preview-05-06"
	case models.PerformReasoning:
		modelName = "gemini-2.5-pro-preview-05-06" // TODO: Reasoning
	default:
		modelName = "gemini-2.0-flash"
	}
	result, err := client.Models.GenerateContent(
		ctx,
		modelName,
		genai.Text(prompt),
		config,
	)
	return result.Text(), err
}
