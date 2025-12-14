package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commits changes and translates the commit message to another language.",
	Long: `This command performs a git commit operation and then translates the commit message to a specified target language.
It's useful for projects requiring multilingual commit logs or for teams working across different linguistic backgrounds.

Example:
$ ct commit "Initial commit" --lang es
$ ct commit "Commit inicial" --lang en`,
	Run: func(cmd *cobra.Command, args []string) {
		lang, _ := cmd.Flags().GetString("lang")
		if len(args) < 1 {
			fmt.Println(`commit message is required. Example: ct commit "Initial commit"`)
			return
		}

		commitMessage := args[0]

		fmt.Println("commit: ", commitMessage)
		fmt.Println("commit called with lang:", lang)
	},
}

func init() {
	commitCmd.Flags().StringP("lang", "l", "en", "Target language for translation")
	rootCmd.AddCommand(commitCmd)
}
