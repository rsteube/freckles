# Add

Copies a dotfile to the repository and replaces it with a symlink.

```go
{{#include ../../../../cmd/freckles/cmd/add.go:cmd}}
```

Completion is provided with [ActionFiles] and a couple of tricks.

```go
{{#include ../../../../cmd/freckles/cmd/add.go:positional}}
```

![](./add/add.cast)

> First of all, dotfiles are located in your home folder.
> For convenience, the workdir in [Context] is changed using [ChdirF] and [`traverse.UserHomeDir`].
>
> ---
>
> Then dotfiles are usually hidden.
> But [ActionFiles] only shows them when the `.` prefix is already present.
> By altering the value in [Context] and **explicitly** invoking [ActionFiles] with it we can force this behaviour.
>
> ---
>
> [Batch] not only enables concurrent invocation of [Actions].
> It also provides a neat way to add them **conditionally**.
> See [ActionRefs] for a complex example.

[Actions]:https://carapace-sh.github.io/carapace/carapace/action.html
[ActionFiles]:https://carapace-sh.github.io/carapace/carapace/defaultActions/actionFiles.html
[ActionRefs]:https://github.com/carapace-sh/carapace-bin/blob/master/pkg/actions/tools/git/ref.go
[Batch]:https://carapace-sh.github.io/carapace/carapace/batch.html
[ChdirF]:https://carapace-sh.github.io/carapace/carapace/action/chdirF.html
[Context]:https://carapace-sh.github.io/carapace/carapace/context.html
[`traverse.UserHomeDir`]:https://pkg.go.dev/github.com/carapace-sh/carapace@v1.8.0/pkg/traverse#UserHomeDir
