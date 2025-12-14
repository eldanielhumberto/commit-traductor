package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push changes to the remote repository",
	Long: `Push changes to the remote repository.
By default, pushes to the current branch on the 'origin' remote.
You can specify a different remote and/or branch using flags.

Examples:
  ct push                           # Push current branch to origin
  ct push --branch main             # Push to main branch
  ct push --remote upstream         # Push to upstream remote
  ct push --branch dev --remote fork # Push to dev branch on fork remote`,
	RunE: func(cmd *cobra.Command, args []string) error {
		remote, _ := cmd.Flags().GetString("remote")
		branch, _ := cmd.Flags().GetString("branch")

		// If no branch specified, get current branch
		if branch == "" {
			currentBranch, err := getCurrentBranch()
			if err != nil {
				return fmt.Errorf("failed to get current branch: %w", err)
			}
			branch = currentBranch
		}

		fmt.Printf("Pushing to %s/%s...\n", remote, branch)

		// Execute git push
		cmdGit := exec.Command("git", "push", remote, branch)
		output, err := cmdGit.CombinedOutput()
		if err != nil {
			return fmt.Errorf("git push failed: %s", string(output))
		}

		fmt.Println(string(output))
		fmt.Printf("✓ Successfully pushed to %s/%s\n", remote, branch)
		return nil
	},
}

// getCurrentBranch returns the name of the current git branch
func getCurrentBranch() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func init() {
	pushCmd.Flags().StringP("remote", "r", "origin", "Remote repository to push to")
	pushCmd.Flags().StringP("branch", "b", "", "Branch to push (defaults to current branch)")
	rootCmd.AddCommand(pushCmd)
}
