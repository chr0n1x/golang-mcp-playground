package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	stdlog "log"
	"os"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	log "github.com/sirupsen/logrus"
)

var startTime = time.Now()

// absolute hack
func stdioLog(msg string) {
	_, _ = fmt.Fprintf(os.Stderr, fmt.Sprintf("%s\n", msg))
}

func helloHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, ok := request.Params.Arguments["name"].(string)
	if !ok {
		return nil, errors.New("name must be a string")
	}

	age := time.Since(startTime).String()

	return mcp.NewToolResultText(
		fmt.Sprintf("Hello, %s! I am %s human time units old", name, age),
	), nil
}

func main() {
	// Create MCP server
	s := server.NewMCPServer(
		"dada demo 🚀",
		"1.0.0",
		server.WithLogging(),
		server.WithPromptCapabilities(true),
	)

	tool := mcp.NewTool("hello_dada",
		mcp.WithDescription("Say hello to my dada"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Name of the person to greet"),
		),
	)
	s.AddTool(tool, helloHandler)

	logger := log.New()
	stdioServer := server.NewStdioServer(s)
	stdLogger := stdlog.New(logger.Writer(), "stdioserver", 0)
	stdioServer.SetErrorLogger(stdLogger)
	ctx := context.Background()

	// Start listening for messages
	errC := make(chan error, 1)
	go func() {
		in, out := io.Reader(os.Stdin), io.Writer(os.Stdout)
		errC <- stdioServer.Listen(ctx, in, out)
	}()

	stdioLog("dada mcp server running in stdio")

	// Wait for shutdown signal
	select {
	case <-ctx.Done():
		logger.Info("shutting down server...")
	case err := <-errC:
		if err != nil {
			logger.Error(fmt.Errorf("error running server: %w", err))
		}
	}
}
