# Golang MCP Template

Barebones stdio MCP server template running in a docker container.

example `oterm` config in `~/.local/share/oterm/config.json`

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

The same `mcpServers` config can be used in CursorAI.

build the image in the command via `docker compose build`.
