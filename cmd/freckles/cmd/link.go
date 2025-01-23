package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/rsteube/freckles/cmd/freckles/cmd/action"
	"github.com/rsteube/freckles/pkg/freckles"
	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "link dotfiles",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		freckles.Walk(func(freckle freckles.Freckle) error {
			if err := freckle.Symlink(false); err != nil {
				println(err.Error())
			}
			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(linkCmd)

	// TODO update command to only link given freckle
	carapace.Gen(linkCmd).PositionalAnyCompletion(
		action.ActionFreckles(),
	)
}
