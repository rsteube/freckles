package cmd

import (
	"fmt"

	"github.com/rsteube/dotfiles-bin/pkg/dotfiles"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list files",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		dotfiles.Walk(func(dotfile dotfiles.Dotfile) error {
			fmt.Println(dotfile.Path)
			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
