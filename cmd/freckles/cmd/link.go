package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/freckles-bin/pkg/dotfiles"
	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "link dotfiles",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		dotfiles.Walk(func(d dotfiles.Dotfile) error {
			if err := d.Symlink(false); err != nil {
				println(err.Error())
			}
			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(linkCmd)

	carapace.Gen(linkCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if dotfiledir, err := c.Abs("~/.local/share/dotfiles"); err != nil {
				return carapace.ActionMessage(err.Error())
			} else {
				return carapace.ActionFiles().Chdir(dotfiledir).Invoke(c).Filter([]string{".git/"}).ToA()
			}
		}),
	)
}
