# Root

[Carapace] is based on [Cobra] which has its own completion logic.

By default [Cobra] adds a `completion` subcommand.
This can be prevented with [`CompletionOptions`].

```go
{{#include ../../../../cmd/freckles/cmd/root.go:cmd}}
```

[Carapace] instead uses a hidden `_carapace` subcommand which is added by [`Gen`].

```go
{{#include ../../../../cmd/freckles/cmd/root.go:init}}
```

Completion scripts are generated with `freckles _carapace [shell]`.
Most of these can be directly sourced.

> [Carapace] has a basic shell detection mechanism, so in most cases `[shell]` is optional.

[Carapace]:https://carapace.sh
[Cobra]:https://github.com/spf13/cobra
[`CompletionOptions`]:https://pkg.go.dev/github.com/spf13/cobra#CompletionOptions
[`Gen`]:https://pkg.go.dev/github.com/carapace-sh/carapace#Gen
