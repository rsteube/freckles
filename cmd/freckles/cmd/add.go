package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/carapace-sh/freckles/pkg/freckles"
	"github.com/spf13/cobra"
)

// ANCHOR: cmd
var addCmd = &cobra.Command{
	Use:   "add [FILE]...",
	Short: "add dotfiles",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			freckle := freckles.Freckle{Path: arg}
			if err := freckle.Add(false); err != nil {
				println(err.Error())
			}
		}
	},
}

// ANCHOR_END: cmd

func init() {
	rootCmd.AddCommand(addCmd)

	// ANCHOR: positional
	carapace.Gen(addCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			batch := carapace.Batch(
				carapace.ActionFiles(),
			)
			if c.Value == "" {
				batch = append(batch, carapace.ActionCallback(func(c carapace.Context) carapace.Action {
					c.Value = "."
					return carapace.ActionFiles().Invoke(c).ToA()
				}))
			}
			return batch.ToA().ChdirF(traverse.UserHomeDir)
		}),
	)
	// ANCHOR: positional
}
