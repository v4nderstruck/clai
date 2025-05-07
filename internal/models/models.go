package models

type ThinkingLevel int

const (
	FastResponse ThinkingLevel = iota
	NormalResponse
	PerformReasoning
)

type ModelFunctions interface {
	ModelHelp() string
	OneShotPrompt(thinkLevel ThinkingLevel, systemPrompt string, prompt string) (string, error)
	// OneShotPrompt(thinkLevel ThinkingLevel, systemPrompt string, prompt string) string
}
