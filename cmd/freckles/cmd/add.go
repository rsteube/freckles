package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/freckles-bin/pkg/freckles"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [FILE]...",
	Short: "add dotfiles",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			d := freckles.Freckle{Path: arg}
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
			if home, err := c.Abs("~/"); err != nil {
				return carapace.ActionMessage(err.Error())
			} else {
				return carapace.ActionFiles().Chdir(home)
			}
		}),
	)
}
