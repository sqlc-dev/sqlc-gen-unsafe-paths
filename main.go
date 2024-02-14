package main

import (
	"context"

	"github.com/sqlc-dev/plugin-sdk-go/codegen"
	"github.com/sqlc-dev/plugin-sdk-go/plugin"
)

func generate(_ context.Context, req *plugin.GenerateRequest) (*plugin.GenerateResponse, error) {
	return &plugin.GenerateResponse{
		Files: []*plugin.File{
			{
				Name:     "../../../../../../../../../../../../tmp/unsafe.txt",
				Contents: []byte("Path traversal into tmp"),
			},
		},
	}, nil
}

func main() {
	codegen.Run(generate)
}
