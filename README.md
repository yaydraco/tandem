<p align="center"><video src="https://github.com/user-attachments/assets/bbfd3fdc-a196-4070-8a7e-c9fad97a322d
" width="800" controls></video></p>

<p>
Swarm of AI Agents to assist in a penetration testing engagement given a RoE.md stating Rules of Engagement.
</p>

## Configuration

### Setting up API Keys

Before using Tandem, you need to configure API keys for the AI providers you want to use. 

1. Copy the example environment file:
   ```shell
   cp .example.env .env
   ```

2. Edit the `.env` file and add your API keys for the providers you plan to use:
   ```
   GEMINI_API_KEY=your_gemini_api_key_here
   OPENAI_API_KEY=your_openai_api_key_here
   GROQ_API_KEY=your_groq_api_key_here
   OPENROUTER_API_KEY=your_openrouter_api_key_here
   VERTEX_API_KEY=your_vertex_api_key_here
   XAI_API_KEY=your_xai_api_key_here
   ANTHROPIC_API_KEY=your_anthropic_api_key_here
   COPILOT_API_KEY=your_copilot_api_key_here
   ```

   **Note**: You don't need to configure all providers - only add keys for the services you want to use.

### Agent Configuration

Tandem's behavior is controlled by the `.tandem/swarm.json` configuration file, which defines the AI agents and their roles. The default configuration includes several specialized agents:

#### Available Providers
The following AI providers are supported:
- **Groq**: Fast inference for various open-source models
- **Anthropic**: Claude models for reasoning and analysis
- **OpenAI**: GPT models for general-purpose tasks
- **Gemini**: Google's AI models
- **OpenRouter**: Access to multiple AI models through a single API
- **Vertex AI**: Google Cloud AI platform
- **GitHub Copilot**: AI pair programming assistant
- **xAI**: Grok and other xAI models

#### Default Agents

**Reconnoiter Agent (Fredrick Russell Burnham)**
- **Role**: OffSec PEN-300 certified penetration tester
- **Purpose**: Handles reconnaissance phase of penetration testing
- **Tools**: Kali Linux CLI tools for reconnaissance
- **Model**: `moonshotai/kimi-k2-instruct`

**Summarizer Agent**
- **Role**: Conversation summarization and context management
- **Purpose**: Provides concise summaries to maintain workflow continuity
- **Model**: `copilot.claude-sonnet-4`

**Title Agent**
- **Role**: Generate concise conversation titles
- **Purpose**: Creates one-liner titles based on user messages
- **Constraint**: Maximum 50 characters

**Orchestrator Agent**
- **Role**: Task coordination and assignment
- **Purpose**: Assigns penetration testing tasks to appropriate specialized agents
- **Model**: `copilot.gemini-2.5-pro`

#### Customizing Agent Configuration

You can modify `.tandem/swarm.json` to:
- Add new agents with custom roles and capabilities
- Modify existing agent instructions and goals
- Change which AI models agents use
- Configure agent-specific tools and permissions
- Adjust debug settings and provider configurations

Example agent structure:
```json
{
  "agents": {
    "your_agent_name": {
      "name": "Display Name",
      "agentId": "unique_id",
      "model": "provider.model-name",
      "description": "What this agent does",
      "goal": "The agent's primary objective",
      "instructions": [
        "Specific instruction 1",
        "Specific instruction 2"
      ],
      "tools": ["tool1", "tool2"]
    }
  }
}
```

## Usage

After configuring your API keys and agent settings:

1. **Set up your engagement context**: Create a `RoE.md` file in your working directory containing the Rules of Engagement for your penetration testing engagement.

2. **Run Tandem**: Start the TUI interface to interact with your AI agent swarm:
   ```shell
   tandem
   ```

3. **Interact with agents**: Use the interface to communicate with specialized agents for different phases of your penetration testing workflow.

## SSH Server Mode

Tandem can run as an SSH server, enabling remote multi-user access to the penetration testing interface. This is useful for distributed teams, remote engagements, or when providing centralized access to tandem's capabilities.

### Quick SSH Demo

To quickly test the SSH functionality:

1. **Set up the demo**:
   ```shell
   ./setup-ssh-demo.sh
   ```

2. **Start the SSH server**:
   ```shell
   ./tandem ssh --host localhost --port 2222
   ```

3. **Connect from another terminal**:
   ```shell
   ssh localhost -p 2222
   ```

You'll see the tandem TUI interface running over SSH, allowing remote access to all penetration testing capabilities.

### SSH Server Configuration

The SSH server supports various configuration options:

```shell
./tandem ssh [flags]

Flags:
  -c, --cwd string        Current working directory
  -d, --debug             Enable debug logging
      --host string       Host to bind the SSH server to (default "localhost")
      --key-path string   Path to SSH host key (will be generated if doesn't exist) (default ".ssh/tandem_host_key")
      --port int          Port to bind the SSH server to (default 2222)
```

### SSH Features

- **Multi-user support**: Multiple users can connect simultaneously, each with their own TUI session
- **Automatic key management**: SSH host keys are automatically generated if they don't exist
- **Session isolation**: Each SSH connection gets its own isolated tandem session
- **Graceful shutdown**: Clean termination of all SSH sessions on server shutdown
- **Secure by default**: Uses standard SSH authentication and encryption

### Production Deployment

For production use:

1. **Configure proper authentication**: Set up SSH keys or other authentication methods as needed
2. **Network security**: Ensure the SSH port is properly firewalled and secured
3. **Monitoring**: Enable logging to monitor SSH connections and usage
4. **Resource management**: Consider resource limits for concurrent sessions

Example production configuration:
```shell
./tandem ssh --host 0.0.0.0 --port 2222 --key-path /etc/tandem/ssh_host_key
```

## Development Instructions
1. This project uses **Nix flake** for setting up a consistent development environment across the team, and we propose you do the same.  
2. Create a .env file before running the ```nix develop``` command. refer to ```.example.env``` to create one.
```shell
# given you have created a .env and nix is present on your sys.
nix develop
```
