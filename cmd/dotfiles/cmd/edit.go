package cmd

import (
	"os"
	"os/exec"

	"github.com/mitchellh/go-homedir"
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [FILE]",
	Short: "edit a dotfile",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if dotfiledir, err := homedir.Expand("~/.local/share/dotfiles"); err == nil {
			c := exec.Command("nvim", dotfiledir+"/"+args[0])
			c.Stdin = os.Stdin
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			c.Run()
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	carapace.Gen(editCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if dotfiledir, err := homedir.Expand("~/.local/share/dotfiles"); err != nil {
				return carapace.ActionMessage(err.Error())
			} else {
				return ActionSubDirectoryFiles(dotfiledir)
			}
		}),
	)
}
