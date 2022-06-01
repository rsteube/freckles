package cmd

import (
	"github.com/rsteube/freckles-bin/pkg/freckles"
	"github.com/spf13/cobra"
)

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "verify symlink status",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		freckles.Walk(func(d freckles.Freckle) error {
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
