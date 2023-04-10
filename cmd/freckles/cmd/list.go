package cmd

import (
	"fmt"
	"os"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/rsteube/freckles/pkg/freckles"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list files",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		freckles.Walk(func(freckle freckles.Freckle) error {
			_style := style.ForPathExt(freckles.FreckleDir()+"/"+freckle.Path, carapace.NewContext(args...))
			fmt.Println(format(freckle.Path, _style))
			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func format(s, _style string) string {
	if fileInfo, _ := os.Stdout.Stat(); (fileInfo.Mode() & os.ModeCharDevice) == 0 {
		return s
	}
	return fmt.Sprintf("\033[%vm%v\033[0m", style.SGR(_style), s)
}
