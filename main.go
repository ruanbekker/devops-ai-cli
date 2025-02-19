package main

import (
	"fmt"
	"github.com/ruanbekker/go-cli-starter/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

