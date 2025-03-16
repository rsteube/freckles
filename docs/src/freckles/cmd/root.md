# Root

[Carapace] uses [Cobra] to define (sub)commands and flags:


```go
{{#include ../../../../cmd/freckles/cmd/root.go:cmd}}
```

> [Cobra] has its own completion logic and adds a `completion` subcommand by default.
> This can be prevented with [`CompletionOptions`].
>
> Historically, [Cobra] had no dynamic completion, and development in this regard was quite stale.
> Without the burden of backward compatibility and the massive amount of [dependents],
> [Carapace] could develop much more quickly and push the boundary of what's possible.
>
> Here's the [cornerstone] that ultimately became [Carapace].

## Gen

[Gen] adds a hidden `_carapace` subcommand which handles [script](#script) generation, [completions](#export), and [macros](#macro).

```go
{{#include ../../../../cmd/freckles/cmd/root.go:gen}}
```

> Calling `Gen` with the root command is enough to add completion for commands and flags.
> But be aware that in [Carapace] argument completion is **explicit**.
> So an undefined completion does not cause an implicit file completion fallback.

## Script

Completion scripts are generated with `freckles _carapace [shell]`.
Most of these can be directly sourced.

![](./root/script.cast)

> [Carapace] has a basic [shell detection] mechanism, so in most cases `[shell]` is optional.

## Export

Shell scripts in [Carapace] are just thin layers to integrate with the corresponding shell.

Aside from shell-specific output:

```sh
freckles _carapace [shell] freckles [ARGS]...
```

[Export] provides a more generic `json` representation of completions:

```sh
freckles _carapace export freckles [ARGS]...
```

![](./root/export.cast)

> [Carapace] (binary) essentially acts as a central registry for all of your completions.
>
> With [#1336] packages will be able to register completions system-wide using [Specs]:
>
> ```yaml
> # yaml-language-server: $schema=https://carapace.sh/schemas/command.json
> name: freckles
> parsing: disabled
> completion:
>   positionalany: ["$carapace.bridge.Carapace([freckles])"]
> ```
>
> This has several benefits:
> - it avoids the [startup delay] issue
> - it provides a central registration point for all shells
> - it enables [embedding] with [`bridge.ActionCarapaceBin`]
> - it uses the newest version of [Carapace] for shell integration

## Macro

[Actions] with the signature of [MacroI], [MacroN], or [MacroV] can be exposed as a [custom macro] for others to consume.

```go
{{#include ../../../../cmd/freckles/cmd/root.go:macro}}
```

![](./root/macro.cast)

> [Macros] provide a way to _loosely_ share [Actions] between applications.
> 
> More on this at [Init#Clone](./init.md#clone) and [Edit#Action](./edit.md#action).

[#1336]:https://github.com/carapace-sh/carapace-bin/issues/1336
[`bridge.ActionCarapaceBin`]:https://pkg.go.dev/github.com/carapace-sh/carapace-bridge/pkg/actions/bridge#ActionCarapaceBin
[Actions]:https://carapace-sh.github.io/carapace/carapace/action.html
[Carapace]:https://carapace.sh
[Cobra]:https://github.com/spf13/cobra
[`CompletionOptions`]:https://pkg.go.dev/github.com/spf13/cobra#CompletionOptions
[cornerstone]:https://github.com/spf13/cobra/pull/646#issuecomment-442267487
[custom macro]:https://carapace-sh.github.io/carapace-spec/carapace-spec/macros/custom.html#custom
[dependents]:https://github.com/spf13/cobra/network/dependents
[embedding]:https://carapace-sh.github.io/carapace-bin/spec/embed.html
[Export]:https://carapace-sh.github.io/carapace/carapace/export.html
[Gen]:https://carapace-sh.github.io/carapace/carapace/gen.html
[MacroI]:https://pkg.go.dev/github.com/carapace-sh/carapace-spec#MacroI
[MacroN]:https://pkg.go.dev/github.com/carapace-sh/carapace-spec#MacroN
[Macros]:https://carapace-sh.github.io/carapace-spec/carapace-spec/macros.html
[MacroV]:https://pkg.go.dev/github.com/carapace-sh/carapace-spec#MacroV
[shell detection]:https://github.com/carapace-sh/carapace/blob/master/pkg/ps/ps.go
[Specs]:https://carapace-sh.github.io/carapace-bin/spec.html
[startup delay]:https://jzelinskie.com/posts/dont-recommend-sourcing-shell-completion/
