package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/dotfiles-bin/pkg/dotfiles"
	"github.com/rsteube/invoke-completion/pkg/invoke"
	"github.com/spf13/cobra"
)

var gitCmd = &cobra.Command{
	Use:                "git",
	Short:              "invoke git on dotfile directory",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command("git", append([]string{"-C", dotfiles.DotfileDir()}, args...)...)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()
	},
}


//go:linkname actionRawValues carapace.actionRawValues 
func actionRawValues(rawValues ...common.RawValue) []string

func init() {
	rootCmd.AddCommand(gitCmd)

	carapace.Gen(gitCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			os.Chdir(dotfiles.DotfileDir()) // TODO support -C in git
			rawValues, err := invoke.InvokeElvish(fmt.Sprintf("git %v", strings.Join(append(c.Args, c.CallbackValue), " ")))
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0)
			for _, rawValue := range rawValues {
				vals = append(vals, strings.TrimSuffix(rawValue.Value, " "), rawValue.Description) // TODO access internal rawValue from carapace
			}
			return carapace.ActionValuesDescribed(vals...)
		}),
	)
}
