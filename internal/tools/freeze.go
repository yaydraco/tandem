package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
)

const (
	FreezeToolName = "freeze"
)

type FreezeArgs struct {
	Input    string `json:"input"`
	Output   string `json:"output,omitempty"`
	Language string `json:"language,omitempty"`
	Theme    string `json:"theme,omitempty"`
	Config   string `json:"config,omitempty"`
}

type Freeze struct{}

func NewFreeze() BaseTool {
	return &Freeze{}
}

func (f *Freeze) Info() ToolInfo {
	return ToolInfo{
		Name:        FreezeToolName,
		Description: "A tool to create beautiful code screenshots using Freeze. Freeze generates images of code with syntax highlighting, themes, and customizable styling options.",
		Parameters: map[string]any{
			"input": map[string]any{
				"type":        "string",
				"description": "Path to the input file or directory containing code to screenshot. Can also be a specific file path.",
			},
			"output": map[string]any{
				"type":        "string",
				"description": "Optional output file path for the generated image. If not specified, Freeze will use default naming.",
			},
			"language": map[string]any{
				"type":        "string",
				"description": "Optional language for syntax highlighting (e.g., 'go', 'python', 'javascript'). If not specified, Freeze will auto-detect.",
			},
			"theme": map[string]any{
				"type":        "string",
				"description": "Optional theme for the code screenshot (e.g., 'dracula', 'github', 'monokai'). Defaults to auto-detect or config setting.",
			},
			"config": map[string]any{
				"type":        "string",
				"description": "Optional path to a Freeze configuration file with custom styling options.",
			},
		},
		Required: []string{"input"},
	}
}

func (f *Freeze) Run(ctx context.Context, call ToolCall) (ToolResponse, error) {
	var args FreezeArgs
	if err := json.Unmarshal([]byte(call.Input), &args); err != nil {
		return NewTextErrorResponse("Failed to parse Freeze arguments: " + err.Error()), nil
	}

	if args.Input == "" {
		return NewTextErrorResponse("input is required for Freeze tool"), nil
	}

	// Check if Freeze is installed
	if err := checkFreezeInstalled(); err != nil {
		return NewTextErrorResponse("Freeze is not installed: " + err.Error()), nil
	}

	// Build Freeze command
	cmdArgs := []string{args.Input}
	
	if args.Output != "" {
		cmdArgs = append(cmdArgs, "--output", args.Output)
	}
	
	if args.Language != "" {
		cmdArgs = append(cmdArgs, "--language", args.Language)
	}
	
	if args.Theme != "" {
		cmdArgs = append(cmdArgs, "--theme", args.Theme)
	}
	
	if args.Config != "" {
		cmdArgs = append(cmdArgs, "--config", args.Config)
	}

	cmd := exec.CommandContext(ctx, "freeze", cmdArgs...)
	
	output, err := cmd.CombinedOutput()
	if err != nil {
		return NewTextErrorResponse(fmt.Sprintf("Freeze command failed: %s\nOutput: %s", err.Error(), string(output))), nil
	}

	return NewTextResponse(fmt.Sprintf("Freeze screenshot created successfully!\nOutput: %s", string(output))), nil
}

func checkFreezeInstalled() error {
	cmd := exec.Command("freeze", "--version")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Freeze not found in PATH. Please install Freeze from https://github.com/charmbracelet/freeze")
	}
	return nil
}