# Go CLI Starter Project

```go
// Project Structure
/*
my-cli/
├── cmd/
│   └── root.go
├── internal/
│   └── handler/
│       └── handler.go
├── pkg/
│   └── utils/
│       └── utils.go
├── main.go
├── go.mod
└── README.md
*/

// go.mod
// Initialize with: go mod init github.com/yourusername/my-cli
/*
module github.com/yourusername/my-cli

go 1.22

require (
	github.com/spf13/cobra v1.8.0
	github.com/spf13/viper v1.18.2
)
*/

// main.go
package main

import (
	"fmt"
	"os"

	"github.com/yourusername/my-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// cmd/root.go
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yourusername/my-cli/internal/handler"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "my-cli",
		Short: "A brief description of your CLI application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application.`,
		Run: func(cmd *cobra.Command, args []string) {
			handler.HandleRootCommand()
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	// Global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.my-cli.yaml)")
	
	// Local flags
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	
	// Example of adding a subcommand
	rootCmd.AddCommand(newExampleCmd())
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".my-cli" (without extension)
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".my-cli")
	}

	// Read in environment variables that match
	viper.AutomaticEnv()

	// If a config file is found, read it in
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// Example of a subcommand
func newExampleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "example",
		Short: "An example of a subcommand",
		Long:  `This is an example of how to create subcommands in your CLI application.`,
		Run: func(cmd *cobra.Command, args []string) {
			verbose, _ := cmd.Flags().GetBool("verbose")
			handler.HandleExampleCommand(verbose)
		},
	}

	// Add flags specific to this command
	cmd.Flags().BoolP("verbose", "v", false, "Enable verbose output")
	
	return cmd
}

// internal/handler/handler.go
package handler

import (
	"fmt"

	"github.com/yourusername/my-cli/pkg/utils"
)

// HandleRootCommand handles the root command
func HandleRootCommand() {
	fmt.Println("Root command executed")
	fmt.Println("Current time:", utils.GetCurrentTime())
}

// HandleExampleCommand handles the example subcommand
func HandleExampleCommand(verbose bool) {
	if verbose {
		fmt.Println("Example command executed with verbose flag")
		fmt.Println("Additional details:", utils.GetSystemInfo())
	} else {
		fmt.Println("Example command executed")
	}
}

// pkg/utils/utils.go
package utils

import (
	"fmt"
	"runtime"
	"time"
)

// GetCurrentTime returns the current time as a formatted string
func GetCurrentTime() string {
	return time.Now().Format(time.RFC3339)
}

// GetSystemInfo returns basic system information
func GetSystemInfo() string {
	return fmt.Sprintf("OS: %s, Architecture: %s, CPUs: %d", 
		runtime.GOOS, 
		runtime.GOARCH, 
		runtime.NumCPU())
}
```
