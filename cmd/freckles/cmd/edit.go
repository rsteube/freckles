package cmd

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/carapace-sh/carapace"
	"github.com/rsteube/freckles/cmd/freckles/cmd/action"
	"github.com/rsteube/freckles/pkg/freckles"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [FILE]",
	Short: "edit a dotfile",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command(editor(), freckles.FreckleDir()+"/"+args[0])
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	carapace.Gen(editCmd).PositionalCompletion(
		action.ActionFreckles(),
	)
}

func editor() string { // source github.com/cli/cli
	defaultEditor := "vi"
	if runtime.GOOS == "windows" {
		defaultEditor = "notepad"
	} else if g := os.Getenv("GIT_EDITOR"); g != "" {
		defaultEditor = g
	} else if v := os.Getenv("VISUAL"); v != "" {
		defaultEditor = v
	} else if e := os.Getenv("EDITOR"); e != "" {
		defaultEditor = e
	}
	return defaultEditor
}
