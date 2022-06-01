package cmd

import (
	"github.com/mitchellh/go-homedir"
	"github.com/rsteube/carapace"
	"github.com/rsteube/dotfiles-bin/pkg/dotfiles"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [FILE]...",
	Short: "add dotfiles",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			d := dotfiles.Dotfile{Path: arg}
			if err := d.Add(false); err != nil {
				println(err.Error())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	carapace.Gen(addCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if home, err := homedir.Dir(); err != nil {
				return carapace.ActionMessage(err.Error())
			} else {
				return carapace.ActionFiles().Chdir(home)
			}
		}),
	)
}
