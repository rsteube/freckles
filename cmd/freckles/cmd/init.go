package cmd

import (
	"os"
	"os/exec"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/rsteube/freckles/pkg/freckles"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init freckles folder",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO handle error when git is missing
		if cmd.Flag("clone").Changed {
			c := exec.Command("git", "clone", cmd.Flag("clone").Value.String(), freckles.FreckleDir())
			c.Stdin = os.Stdin
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			c.Run()
		} else {
			c := exec.Command("git", "init", freckles.FreckleDir())
			c.Stdin = os.Stdin
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			c.Run()
		}

		if _, err := os.Stat(freckles.FreckleDir() + ".frecklesignore"); os.IsNotExist(err) {
			if err := os.WriteFile(freckles.FreckleDir()+".frecklesignore", []byte(".git\n.frecklesignore\n"), os.ModePerm); err != nil {
				panic(err.Error())
			}
		}
	},
}

func init() {
	initCmd.Flags().String("clone", "", "clone existing repo")

	rootCmd.AddCommand(initCmd)

	carapace.Gen(initCmd).FlagCompletion(carapace.ActionMap{
		"clone": git.ActionRepositorySearch(git.SearchOpts{}.Default()),
	})
}
