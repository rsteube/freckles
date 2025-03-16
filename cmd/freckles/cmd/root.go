package cmd

import (
	"github.com/carapace-sh/carapace"
	spec "github.com/carapace-sh/carapace-spec"
	"github.com/carapace-sh/freckles/cmd/freckles/cmd/action"
	"github.com/spf13/cobra"
)

// ANCHOR: cmd
var rootCmd = &cobra.Command{
	Use:   "freckles",
	Short: "A simple dotfile manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

// ANCHOR_END: cmd

func Execute(version string) error {
	rootCmd.Version = version
	return rootCmd.Execute()
}

func init() {
	// ANCHOR: gen
	carapace.Gen(rootCmd)
	// ANCHOR_END: gen

	// ANCHOR: macro
	spec.AddMacro("freckles", spec.MacroN(action.ActionFreckles))
	spec.Register(rootCmd)
	// ANCHOR_END: macro
}
