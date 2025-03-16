# Init

Creates a new [Git] repository at `~/.local/share/freckles`.

```go
{{#include ../../../../cmd/freckles/cmd/init.go:cmd}}
```

![](./init/init.cast)

## Clone

Alternatively an existing remote repository can be cloned with the `--clone` flag.

Completion is provided by [ActionRepositorySearch].

```go
{{#include ../../../../cmd/freckles/cmd/init.go:flagcompletion}}
```

![](./init/clone.cast)

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
>
> Note that for performance reasons only the first 100 search results are presented.
> Fast response times are important which limits what can be done in [Carapace].

[ActionRepositorySearch]:https://pkg.go.dev/github.com/carapace-sh/carapace-bin/pkg/actions/tools/git#ActionRepositorySearch
[Actions]:https://carapace-sh.github.io/carapace/carapace/action.html
[Carapace]:https://carapace.sh
[Git]:https://git-scm.com/
