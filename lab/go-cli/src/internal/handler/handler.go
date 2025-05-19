package handler

import (
	"fmt"

	"github.com/JaimeStill/go-cli/pkg/utils"
)

func HandleRootCommand() {
	fmt.Println("Root command executed")
	fmt.Println("Current time:", utils.GetCurrentTime())
}

func HandleExampleCommand(verbose bool) {
	if verbose {
		fmt.Println("Example command executed with verbose flag")
		fmt.Println("Additional details:", utils.GetSystemInfo())
	} else {
		fmt.Println("Example command executed")
	}
}
