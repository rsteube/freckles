package cmd

import (
	"io/ioutil"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/rsteube/carapace"
	"github.com/rsteube/dotfiles-bin/pkg/dotfiles"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [FILE]...",
	Short: "add dotfiles",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			d := dotfiles.Dotfile{Path: arg}
            if err := d.Add(false); err != nil {
              println(err.Error())
            }
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	carapace.Gen(addCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if home, err := homedir.Dir(); err != nil {
				return carapace.ActionMessage(err.Error())
			} else {
				return ActionSubDirectoryFiles(home)
			}
		}),
	)
}

// ActionSubDirectories completes subdirectories of a given path
//   subdir/subsubdir
//   subdir/subsubder2
func ActionSubDirectoryFiles(path string) carapace.Action {
	return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
		if files, err := ioutil.ReadDir(path + "/" + strings.Join(c.Parts, "/") + "/"); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			dirs := make([]string, 0)
			for _, file := range files {
				if file.IsDir() {
					dirs = append(dirs, file.Name()+"/")
				} else {
					dirs = append(dirs, file.Name())
				}
			}
			return carapace.ActionValues(dirs...)
		}
	})
}
