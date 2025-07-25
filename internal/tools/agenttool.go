package tools

import (
	"context"

	"github.com/Drax-1/tandem/internal/session"
)

const (
	AgentToolName = "agent_tool"
)

type AgentToolParams struct {
	Prompt         string `json:"prompt"`
	AgentID        string `json:"agent_id,omitempty"`
	ExpectedOutput string `json:"expected_output,omitempty"`
}

type AgentTool struct {
	sessions session.Service
}

func (a *AgentTool) Info() ToolInfo {
	return ToolInfo{
		Name:        AgentToolName,
		Description: "A tool to assign penetrationg testing engagement related tasks to agents as per their area of expertise and purview.",
		Parameters: map[string]any{
			"prompt": map[string]any{
				"type":        "string",
				"description": "Prompt to send to the agent",
			},
			"agent_id": map[string]any{
				"type":        "string",
				"description": "ID of the agent to call",
			},
			"expected_output": map[string]any{
				"type":        "string",
				"description": "a json string representing the expected output schema.",
			},
		},
		Required: []string{"prompt", "agent_id", "expected_output"},
	}
}

func (a *AgentTool) Run(ctx context.Context, call ToolCall) (ToolResponse, error) {
	/*
		STEPS:
		- based on the AgentID
	*/

}
