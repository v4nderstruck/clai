package internal

import (
	"fmt"

	"github.com/v4nderstruck/clai/internal/models"
	"github.com/v4nderstruck/clai/internal/models/gemini"
)

type ClaiTool struct {
	AiModel models.ModelFunctions
}

var supportedModels = map[string]models.ModelFunctions{
	"gemini": gemini.NewGeminiModel(),
}

func PrintSupportedModelFamilies() {}



func NewClaiTool(model_name string) (*ClaiTool, error) {
	model, ok := supportedModels[model_name]
	if !ok {
		return nil, fmt.Errorf("selected unsupported model '%s'", model_name)
	}
	c := ClaiTool{
		AiModel: model,
	}
	return &c, nil
}
