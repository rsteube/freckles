package action

import (
	"path/filepath"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/rsteube/freckles/pkg/freckles"
)

func ActionFreckles() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		vals := make([]string, 0)
		freckles.Walk(func(freckle freckles.Freckle) error {
			vals = append(vals, freckle.Path)
			return nil
		})
		return carapace.ActionValues(vals...).MultiParts("/").StyleF(func(s string, sc style.Context) string {
			return style.ForPath(filepath.Join(freckles.Dir(), s), sc)
		})
	})
}
