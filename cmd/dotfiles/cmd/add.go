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
	Use:   "add",
	Short: "add a file",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		d := dotfiles.Dotfile{Path: args[0]}
		if err = d.Add(false); err == nil {
			err = d.Symlink(false)
		}
		return
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	carapace.Gen(addCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(args []string) carapace.Action {
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
	return carapace.ActionMultiParts("/", func(args, parts []string) carapace.Action {
		if files, err := ioutil.ReadDir(path + "/" + strings.Join(parts, "/") + "/"); err != nil {
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
