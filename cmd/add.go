package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add all changes to the staging area",
	Long:  `Add all changes to the staging area`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdGit := exec.Command("git", "add", "-A")
		output, err := cmdGit.CombinedOutput()
		if err != nil {
			return fmt.Errorf("git add failed: %s", string(output))
		}

		fmt.Println("Changes added to staging area")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
