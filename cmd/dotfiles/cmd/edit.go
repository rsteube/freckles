package cmd

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/rsteube/carapace"
	"github.com/rsteube/dotfiles-bin/cmd/dotfiles/cmd/action"
	"github.com/rsteube/dotfiles-bin/pkg/dotfiles"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [FILE]",
	Short: "edit a dotfile",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command(editor(), dotfiles.DotfileDir()+"/"+args[0])
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	carapace.Gen(editCmd).PositionalCompletion(
		action.ActionDotfiles(),
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
