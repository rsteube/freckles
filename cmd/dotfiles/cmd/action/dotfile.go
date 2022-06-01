package action

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/rsteube/dotfiles-bin/pkg/dotfiles"
)

func ActionDotfiles() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		vals := make([]string, 0)
		dotfiles.Walk(func(dotfile dotfiles.Dotfile) error {
			vals = append(vals, dotfile.Path)
			return nil
		})
		return carapace.ActionValues(vals...).Invoke(c).ToMultiPartsA("/").StyleF(func(s string) string {
			return style.ForPath(dotfiles.DotfileDir() + s)
		})
	})
}
