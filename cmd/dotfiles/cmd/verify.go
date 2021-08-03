package cmd

import (
	"github.com/rsteube/dotfiles-bin/pkg/dotfiles"
	"github.com/spf13/cobra"
)

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "verify symlink status",
	Args:  cobra.NoArgs,
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
}
