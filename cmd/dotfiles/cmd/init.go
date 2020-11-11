package cmd

import (
	"os"
	"os/exec"

	"github.com/rsteube/dotfiles-bin/pkg/dotfiles"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init dotfiles folder",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flag("clone").Changed {
			c := exec.Command("git", "clone", cmd.Flag("clone").Value.String(), dotfiles.DotfileDir())
			c.Stdin = os.Stdin
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			c.Run()
		} else {
			c := exec.Command("git", "init", dotfiles.DotfileDir())
			c.Stdin = os.Stdin
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			c.Run()
		}
	},
}

func init() {
	initCmd.Flags().String("clone", "", "clone existing repo")

	rootCmd.AddCommand(initCmd)
}
