package cmd

import (
	"os"
	"os/exec"

	"github.com/rsteube/carapace"
	"github.com/rsteube/dotfiles-bin/cmd/dotfiles/cmd/action"
	"github.com/rsteube/dotfiles-bin/pkg/dotfiles"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [FILE]",
	Short: "edit a dotfile",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command("nvim", dotfiles.DotfileDir()+"/"+args[0])
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	carapace.Gen(editCmd).PositionalCompletion(
		action.ActionDotfiles(),
	)
}
