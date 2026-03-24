package mcpserver

import (
	"encoding/json"

	mcp "trpc.group/trpc-go/trpc-mcp-go"
)

func jsonResult(v any) (*mcp.CallToolResult, error) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return nil, err
	}
	return mcp.NewTextResult(string(data)), nil
}

func errResult(err error) (*mcp.CallToolResult, error) {
	return mcp.NewErrorResult(err.Error()), nil
}

func stringArg(args map[string]any, key string) string {
	if v, ok := args[key].(string); ok {
		return v
	}
	return ""
}
