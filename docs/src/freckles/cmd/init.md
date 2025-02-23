# Init

[Freckles] is backed by a [Git] repository located at `~/.local/share/freckles`.

```go
{{#include ../../../../cmd/freckles/cmd/init.go:command}}
```

![](./init/init.cast)

## Clone

Completion for the `--clone` flag is provided by [ActionRepositorySearch].

```go
{{#include ../../../../cmd/freckles/cmd/init.go:flagcompletion}}
```

![](./init/clone.cast)

> [Macros] provide a way to _loosely_ share [Actions] between applications.
>
> Here, `carapace` is invoked with the macro `tools.git.RepositorySearch` and the current value:
> ```sh
> carapace _carapace macro tools.git.RepositorySearch https://github.com/rsteube/do
> ```
> Which then returns the completion in the [Export](https://carapace-sh.github.io/carapace/carapace/export.html) format.
> ```json
> {
>   "version": "v1.8.0",
>   "messages": [],
>   "nospace": "/",
>   "usage": "",
>   "values": [
>     {
>       "value": "https://github.com/rsteube/docker-mdbook",
>       "display": "docker-mdbook",
>       "description": "mdbook mermaid "
>     },
>     {
>       "value": "https://github.com/rsteube/docker-mdbook-dtmo",
>       "display": "docker-mdbook-dtmo",
>       "description": "mdbook mdbook-mermaid mdbook-toc"
>     },
>     {
>       "value": "https://github.com/rsteube/dotfiles",
>       "display": "dotfiles",
>       "style": "red"
>     }
>   ]
> }
> ```
> ![](./init/macro.cast)

[Actions]:https://carapace-sh.github.io/carapace/carapace/action.html
[ActionRepositorySearch]:https://pkg.go.dev/github.com/carapace-sh/carapace-bin/pkg/actions/tools/git#ActionRepositorySearch
[Freckles]:https://github.com/rsteube/freckles
[Git]:https://git-scm.com/
[Macros]:https://carapace-sh.github.io/carapace-spec/carapace-spec/macros.html
