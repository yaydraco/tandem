package agent

import "github.com/Drax-1/tandem/internal/tools"

// NOTE: concatenate it with the role specific tools
var PenetrationTestingAgentTools = []tools.BaseTool{
	tools.NewDockerCli(),
}

func OrchestratorAgentTools() []tools.BaseTool {
	return []tools.BaseTool{
		// tools.NewAgentTool(),
	}
}
