# Root

[Carapace] is based on [Cobra] and uses it to define (sub)commands and flags:


```go
{{#include ../../../../cmd/freckles/cmd/root.go:cmd}}
```

> [Cobra] has its own completion logic and adds a `completion` subcommand by default.
> This can be prevented with [`CompletionOptions`].

## Gen

[Carapace] instead uses a hidden `_carapace` subcommand which is added by [`Gen`].

```go
{{#include ../../../../cmd/freckles/cmd/root.go:gen}}
```

> Calling `Gen` with the root command is enough to add completion for commands and flags.
> But be aware that in [Carapace] argument completion is **explicit**.
> So an undefined completion does not cause a file completion fallback.

## Script

Completion scripts are generated with `freckles _carapace [shell]`.
Most of these can be directly sourced.

![](./root/script.cast)

> [Carapace] has a basic shell detection mechanism, so in most cases `[shell]` is optional.

## Macro

Actions with the signature of [MacroI], [MacroN], or [MacroV] can be exposed as [custom macro] for others to consume
(see [Clone](./init.md#clone)).

```go
{{#include ../../../../cmd/freckles/cmd/root.go:macro}}
```

![](./root/macro.cast)

[Carapace]:https://carapace.sh
[Cobra]:https://github.com/spf13/cobra
[`CompletionOptions`]:https://pkg.go.dev/github.com/spf13/cobra#CompletionOptions
[`Gen`]:https://pkg.go.dev/github.com/carapace-sh/carapace#Gen

[custom macro]:https://carapace-sh.github.io/carapace-spec/carapace-spec/macros/custom.html#custom

[MacroI]:https://pkg.go.dev/github.com/carapace-sh/carapace-spec#MacroI
[MacroN]:https://pkg.go.dev/github.com/carapace-sh/carapace-spec#MacroN
[MacroV]:https://pkg.go.dev/github.com/carapace-sh/carapace-spec#MacroV
