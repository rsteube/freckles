package cmd

import (
	"github.com/mitchellh/go-homedir"
	"github.com/rsteube/carapace"
	"github.com/rsteube/dotfiles-bin/pkg/dotfiles"
	"github.com/spf13/cobra"
)

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "verify symlink status",
	Run: func(cmd *cobra.Command, args []string) {
        dotfiles.Walk(func(d dotfiles.Dotfile) error {
          if d.Verify() {
              println("[OK]  " + d.Path)
          } else {
              println("[ERR] " + d.Path)
          }
          return nil
        })
	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)

	carapace.Gen(verifyCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if home, err := homedir.Dir(); err != nil {
				return carapace.ActionMessage(err.Error())
			} else {
				return ActionSubDirectoryFiles(home)
			}
		}),
	)
}
