package main

import (
	"os"

	"github.com/koki-develop/gh-x/internal/cmd"
)

func main() {
	if err := cmd.Run(); err != nil {
		os.Exit(1)
	}
}
