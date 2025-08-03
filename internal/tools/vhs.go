package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
)

const (
	VHSToolName = "vhs"
)

type VHSArgs struct {
	TapeFile string `json:"tape_file"`
	Output   string `json:"output,omitempty"`
}

type VHS struct{}

func NewVHS() BaseTool {
	return &VHS{}
}

func (v *VHS) Info() ToolInfo {
	return ToolInfo{
		Name:        VHSToolName,
		Description: "A tool to create terminal recordings and GIFs using VHS (Video Home System). VHS uses tape files to script terminal sessions and generate beautiful terminal recordings.",
		Parameters: map[string]any{
			"tape_file": map[string]any{
				"type":        "string",
				"description": "Path to the VHS tape file (.tape) that contains the recording script. The tape file defines the commands to run, typing speed, and recording settings.",
			},
			"output": map[string]any{
				"type":        "string",
				"description": "Optional output file path for the generated recording. If not specified, VHS will use the output setting from the tape file.",
			},
		},
		Required: []string{"tape_file"},
	}
}

func (v *VHS) Run(ctx context.Context, call ToolCall) (ToolResponse, error) {
	var args VHSArgs
	if err := json.Unmarshal([]byte(call.Input), &args); err != nil {
		return NewTextErrorResponse("Failed to parse VHS arguments: " + err.Error()), nil
	}

	if args.TapeFile == "" {
		return NewTextErrorResponse("tape_file is required for VHS tool"), nil
	}

	// Check if VHS is installed
	if err := checkVHSInstalled(); err != nil {
		return NewTextErrorResponse("VHS is not installed: " + err.Error()), nil
	}

	// Build VHS command
	cmdArgs := []string{args.TapeFile}
	if args.Output != "" {
		cmdArgs = append([]string{"--output", args.Output}, cmdArgs...)
	}

	cmd := exec.CommandContext(ctx, "vhs", cmdArgs...)
	
	output, err := cmd.CombinedOutput()
	if err != nil {
		return NewTextErrorResponse(fmt.Sprintf("VHS command failed: %s\nOutput: %s", err.Error(), string(output))), nil
	}

	return NewTextResponse(fmt.Sprintf("VHS recording created successfully!\nOutput: %s", string(output))), nil
}

func checkVHSInstalled() error {
	cmd := exec.Command("vhs", "--version")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("VHS not found in PATH. Please install VHS from https://github.com/charmbracelet/vhs")
	}
	return nil
}