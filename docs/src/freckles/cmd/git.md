# Git

[Freckles] embeds [Git] as subcommand to manage the repository.

By passing `-C <dir>` it acts as if called from within the repository at `~/.local/share/freckles`.
And with `DisableFlagParsing: true` every argument is seen as positional and passed along.


```go
{{#include ../../../../cmd/freckles/cmd/git.go:cmd}}
```

Completion is provided with [`ActionCarapaceBin`].

```go
{{#include ../../../../cmd/freckles/cmd/git.go:positional}}
```

![](./git/git.cast)

[`ActionCarapaceBin`]:https://pkg.go.dev/github.com/carapace-sh/carapace-bridge/pkg/actions/bridge#ActionCarapaceBin
[Freckles]:https://github.com/rsteube/freckles
[Git]:https://git-scm.com/
