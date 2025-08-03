package main

import (
	"fmt"
	"kong-kli/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Printf("Unable to launch kong-kli")
	}
}
