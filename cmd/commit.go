package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/eldanielhumberto/ct/internal/config"
	"github.com/eldanielhumberto/ct/internal/translator"
	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit [message]",
	Short: "Commits changes and translates the commit message to another language",
	Long: `This command performs a git commit operation and then translates the commit message to a specified target language.
It's useful for projects requiring multilingual commit logs or for teams working across different linguistic backgrounds.

Example:
  ct commit "Initial commit" --lang es
  ct commit "Commit inicial" --lang en`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		commitMessage := args[0]
		lang, _ := cmd.Flags().GetString("lang")

		fmt.Printf("Original message: %s\n", commitMessage)

		// Load configuration
		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("configuration error: %w", err)
		}

		// Create context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// Create translator
		trans, err := translator.New(ctx, cfg.APIKey)
		if err != nil {
			return fmt.Errorf("failed to initialize translator: %w", err)
		}
		defer trans.Close()

		// Translate message
		translated, err := trans.Translate(ctx, commitMessage, lang)
		if err != nil {
			return fmt.Errorf("translation failed: %w", err)
		}

		fmt.Printf("Translated message (%s): %s\n", lang, translated)
		return nil
	},
}

func init() {
	commitCmd.Flags().StringP("lang", "l", "en", "Target language for translation")
	rootCmd.AddCommand(commitCmd)
}
