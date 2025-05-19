package cmd

import (
	"fmt"
	"os"

	"github.com/JaimeStill/go-cli/internal/handler"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "go-cli",
		Short: "A brief description of your CLI application",
		Long: `A longer desciprion that spans muliple lines and likely contains
		examples and usage of using your application.`,
		Run: func(cmd *cobra.Command, args []string) {
			handler.HandleRootCommand()
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(
		&cfgFile,
		"config",
		"",
		"config file (default is $HOME/.go-cli.yaml)",
	)

	rootCmd.Flags().BoolP(
		"toggle",
		"t",
		false,
		"Help message for toggle",
	)

	rootCmd.AddCommand(newExampleCmd())
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".my-cli")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

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

	cmd.Flags().BoolP(
		"verbose",
		"v",
		false,
		"Enable verbose output",
	)

	return cmd
}
