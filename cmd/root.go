package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ct",
	Short: "Translate and standardize git commit messages from the CLI",
	Long: `ct (commit translator) is a command-line tool that helps you write
clean, consistent and readable git commit messages — without thinking
too much about the language.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
