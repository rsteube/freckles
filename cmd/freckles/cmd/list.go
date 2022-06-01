package cmd

import (
	"fmt"

	"github.com/rsteube/carapace/pkg/style"
	"github.com/rsteube/freckles-bin/pkg/dotfiles"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list files",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		dotfiles.Walk(func(dotfile dotfiles.Dotfile) error {
			_style := style.ForPathExt(dotfiles.DotfileDir() + "/" + dotfile.Path)
			fmt.Println(format(dotfile.Path, _style))
			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func format(s, _style string) string {
	return fmt.Sprintf("\033[%vm%v\033[0m", style.SGR(_style), s)
}
