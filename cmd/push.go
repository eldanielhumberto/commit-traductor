package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push changes to the remote repository",
	Long:  `Push changes to the remote repository`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdGit := exec.Command("git", "push")
		output, err := cmdGit.CombinedOutput()
		if err != nil {
			return fmt.Errorf("git push failed: %s", string(output))
		}

		fmt.Println(string(output))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}
