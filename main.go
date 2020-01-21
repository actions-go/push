package main

import (
	"os"

	"github.com/actions-go/push/pkg/push"
	"github.com/actions-go/toolkit/core"
)

func main() {
	if err := push.Push("."); err != nil {
		core.Errorf("failed to push changes: %v", err)
		os.Exit(1)
	}
}
