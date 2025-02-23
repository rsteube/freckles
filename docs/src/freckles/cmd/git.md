# Git

[Freckles] embeds [Git] as subcommand.

By passing `-C <dir>` it acts as if called from within the repository at `~/.local/share/freckles`.
And with `DisableFlagParsing: true` every argument is seen as positional.


```go
{{#include ../../../../cmd/freckles/cmd/git.go:cmd}}
```

```go
{{#include ../../../../cmd/freckles/cmd/git.go:positional}}
```


[Freckles]:https://github.com/rsteube/freckles
[Git]:https://git-scm.com/
