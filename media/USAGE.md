# Media Agent Usage Guide

The Media Agent is a specialized AI agent that creates promotional content for Tandem using VHS (terminal recordings) and Freeze (code screenshots).

## Quick Start

### 1. Using CLI (Non-Interactive Mode)

```bash
# Generate media content with a specific prompt
./tandem -p "Create a terminal recording showing Tandem's help command and version, then take screenshots of the main.go and agent configuration files" -f json

# Request specific media type
./tandem -p "As the media agent, create a VHS recording using the tape file at media/tandem-showcase.tape"

./tandem -p "As the media agent, take a beautiful screenshot of the swarm.json configuration file using Freeze"
```

### 2. Using Interactive Mode

```bash
./tandem
# Then in the TUI, start a conversation and mention the media agent:
# "Can you use the media agent to create promotional content for our latest features?"
```

### 3. Via GitHub Actions

The media agent automatically runs when you push to main or feature branches. The workflow:
- Installs VHS and Freeze
- Builds Tandem
- Generates recordings and screenshots
- Uploads artifacts to the PR/commit

## Tool Parameters

### VHS Tool
- `tape_file`: Path to the .tape file with recording script
- `output`: Optional output file path

### Freeze Tool  
- `input`: Path to code file or directory to screenshot
- `output`: Optional output path for the image
- `language`: Optional language for syntax highlighting
- `theme`: Optional theme (e.g., 'dracula', 'github')
- `config`: Optional config file path

## Examples

### Creating a VHS Recording

The media agent can create terminal recordings by using VHS tape files:

```json
{
  "tape_file": "media/tandem-showcase.tape",
  "output": "media/demo.gif"
}
```

### Taking Code Screenshots

The media agent can create beautiful code screenshots:

```json
{
  "input": "main.go",
  "output": "media/main-screenshot.png", 
  "language": "go",
  "theme": "dracula",
  "config": "media/freeze-config.json"
}
```

## Tips

- Keep recordings under 90 seconds for optimal engagement
- Use the Dracula theme for consistency with Tandem's branding
- Include real usage scenarios in recordings
- Screenshots should highlight key features or changes
- Always test tape files before production use