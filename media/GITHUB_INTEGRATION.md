# GitHub Integration for Media Agent

This document outlines how the Media Agent integrates with GitHub features for automated content generation.

## GitHub Actions Integration

The Media Agent leverages GitHub Actions through the `.github/workflows/generate-media.yml` workflow to:

- **Automatic Triggers**: Runs on pushes to main, develop, feature/*, and release/* branches
- **Manual Triggers**: Supports workflow_dispatch for on-demand media generation
- **PR Integration**: Comments on pull requests with generated media content

## GitHub Copilot Features Used

The Media Agent utilizes GitHub Copilot capabilities in several ways:

### 1. Copilot Model Integration
- Uses `copilot.claude-sonnet-4` for optimal creative and technical reasoning
- Leverages Copilot's code understanding for better screenshot composition
- Benefits from Copilot's context awareness for relevant content generation

### 2. Copilot-Style Instructions
The agent is configured with instructions that mirror Copilot's coding assistant patterns:
- Context-aware content creation
- Focus on demonstrating actual functionality
- Emphasis on developer-focused scenarios

### 3. Integration with GitHub Ecosystem
- Seamless operation within GitHub repositories
- Artifact generation and storage through GitHub Actions
- PR commenting and interaction capabilities

## Secrets Configuration

The workflow requires these GitHub secrets to be configured:

```
GEMINI_API_KEY - For Gemini model access
GROQ_API_KEY - For Groq model access  
OPENAI_API_KEY - For OpenAI model access
ANTHROPIC_API_KEY - For Anthropic/Claude model access
COPILOT_API_KEY - For GitHub Copilot model access
```

## Customization Options

### Workflow Inputs
- `media_type`: Choose between 'both', 'recording', or 'screenshot'
- Triggered automatically on branch pushes
- Manual execution via workflow_dispatch

### Content Types
- **Terminal Recordings**: Showcase CLI interactions and TUI features
- **Code Screenshots**: Highlight key code sections and configurations
- **Combined Content**: Both recordings and screenshots for comprehensive demos

## Usage in Marketing

The generated content can be used for:
- Social media posts (Twitter, LinkedIn)
- Documentation updates
- Release announcements  
- Feature demonstrations
- Onboarding materials

## Best Practices

1. **Recording Length**: Keep under 90 seconds for optimal engagement
2. **Consistency**: Use Dracula theme for brand alignment
3. **Content Quality**: Focus on real usage scenarios
4. **Automation**: Let the workflow handle generation on important pushes
5. **Review**: Always review generated content before external use