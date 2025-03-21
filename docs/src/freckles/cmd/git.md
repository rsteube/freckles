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

> Without arguments [`ActionCarapaceBin`] completes all registered commands.
>
> Here, we limit it to `git`.
> But it could be even more specific by passing subcommands, flags and arguments.
>
> ---
>
> Just as above we change the directory with [Chdir] to the repository.
>
> Note that [Carapace] also supports generic modifications like `-C <dir>` with [PreInvoke].
>
> ---
>
> When fully embedding a command `DisableFlagParsing: true` is the right aproach.
>
> But for [sudo]-like behaviour where the local command has flags as well there is [SetInterspersed].
> Disabling it stops flag parsing after the first positional argument.

[`ActionCarapaceBin`]:https://pkg.go.dev/github.com/carapace-sh/carapace-bridge/pkg/actions/bridge#ActionCarapaceBin
[Carapace]:https://carapace.sh
[Chdir]:https://carapace-sh.github.io/carapace/carapace/action/chdir.html
[Freckles]:https://github.com/carapace-sh/freckles
[Git]:https://git-scm.com/
[sudo]:https://github.com/carapace-sh/carapace-bin/blob/master/completers/sudo_completer/cmd/root.go
[PreInvoke]:https://carapace-sh.github.io/carapace/carapace/gen/preInvoke.html
[SetInterspersed]:https://pkg.go.dev/github.com/spf13/pflag#SetInterspersed
