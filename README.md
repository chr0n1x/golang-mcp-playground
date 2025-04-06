# Golang MCP Playground

Screwing around with MCP, ollama, via oterm

example oterm config in `~/.local/share/oterm/config.json`

```json
{
    "mcpServers": {
        "playground": {
            "args": [
                "run",
                "--rm",
                "-i",
                "golang-mcp-playground-mcp"
            ],
            "command": "docker"
        }
    },
    "splash-screen": true,
    "theme": "textual-dark"
}
```

build the image via `docker compose build`.
