package action

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/rsteube/freckles-bin/pkg/freckles"
)

func ActionFreckles() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		vals := make([]string, 0)
		freckles.Walk(func(dotfile freckles.Freckle) error {
			vals = append(vals, dotfile.Path)
			return nil
		})
		return carapace.ActionValues(vals...).Invoke(c).ToMultiPartsA("/").StyleF(func(s string) string {
			return style.ForPath(freckles.FreckleDir() + s)
		})
	})
}
