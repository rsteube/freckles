# Edit

Opens a dotfile in your editor.

```go
{{#include ../../../../cmd/freckles/cmd/edit.go:cmd}}
```

Completion is provided with a [custom action].

```go
{{#include ../../../../cmd/freckles/cmd/edit.go:positional}}
```

![](./edit/edit.cast)

## Action

[Custom Actions] are simply functions returning [`Action`].

```go
{{#include ../../../../cmd/freckles/cmd/action/freckle.go}}
```

> There are **two main phases** in [Carapace]:
> 1. The creation of the command structure and registration of completions.
> 2. The parsing of the command line and invocation of the corresponding [Action].
>
> ```go
> carapace.Gen(initCmd).PositionalCompletion(
> 	carapace.ActionValues(initCmd.Flag("clone").Value.String()), // (1) completes the default value
> 	carapace.ActionCallback(func(c carapace.Context) carapace.Action {
> 		return carapace.ActionValues(initCmd.Flag("clone").Value.String()) // (2) completes the parsed value
> 	}),
> )
> ```
>
> [Custom Actions] should almost always be wrapped in [ActionCallback] so code is only executed when invoked.
> The [Default Actions] do this implicitly.
>
> ---
>
> The [Walk] function returns dotfiles with their full path within the repository: `path/to/freckle`.
> By modifying the [Action] with [MultiParts] the segments get completed separately.
>
> ---
>
> Additionally, [`style.ForPath`] highlights them with the style defined by the `LS_COLORS` environment variable.


[Action]:https://carapace-sh.github.io/carapace/carapace/action.html
[`Action`]:https://pkg.go.dev/github.com/carapace-sh/carapace#Action
[ActionCallback]:https://carapace-sh.github.io/carapace/carapace/defaultActions/actionCallback.html
[Carapace]:https://carapace.sh
[Custom Actions]:https://carapace-sh.github.io/carapace/carapace/customActions.html
[custom action]:https://carapace-sh.github.io/carapace/carapace/customActions.html
[Default Actions]:https://carapace-sh.github.io/carapace/carapace/defaultActions.html
[MultiParts]:https://carapace-sh.github.io/carapace/carapace/action/multiParts.html
[`style.ForPath`]:https://pkg.go.dev/github.com/carapace-sh/carapace/pkg/style#ForPath
[Walk]:https://pkg.go.dev/github.com/rsteube/freckles/pkg/freckles#Walk
