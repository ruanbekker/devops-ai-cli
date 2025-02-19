package main

import (
	"fmt"
	"github.com/ruanbekker/devops-ai-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

