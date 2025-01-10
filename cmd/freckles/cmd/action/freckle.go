package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/rsteube/freckles/pkg/freckles"
)

func ActionFreckles() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		vals := make([]string, 0)
		freckles.Walk(func(dotfile freckles.Freckle) error {
			vals = append(vals, dotfile.Path)
			return nil
		})
		return carapace.ActionValues(vals...).Invoke(c).ToMultiPartsA("/").StyleF(func(s string, sc style.Context) string {
			return style.ForPath(freckles.FreckleDir()+s, sc)
		})
	})
}
