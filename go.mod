module github.com/rsteube/dotfiles-bin

go 1.12

require (
	github.com/go-git/go-billy/v5 v5.0.0
	github.com/go-git/go-git/v5 v5.2.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/rsteube/carapace v0.13.0
	github.com/rsteube/carapace-bin v0.8.12
	github.com/spf13/cobra v1.3.0
)

replace github.com/spf13/pflag => github.com/cornfeedhobo/pflag v1.1.0
