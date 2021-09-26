// +build tools
// This file is to prevent `go mod tiny` to remove our tools

package tools

import (
	_ "github.com/cortesi/modd/cmd/modd"
	_ "golang.org/x/tools/cmd/goimports"
	_ "gotest.tools/gotestsum"
)
