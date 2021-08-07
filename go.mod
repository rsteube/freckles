module github.com/rsteube/dotfiles-bin

go 1.12

require (
	github.com/go-git/go-billy/v5 v5.0.0
	github.com/go-git/go-git/v5 v5.2.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/rsteube/carapace v0.7.2
	github.com/rsteube/invoke-completion v0.0.0-20210807171303-62f5fcac1c71
	github.com/spf13/cobra v1.2.1
	golang.org/x/sys v0.0.0-20210806184541-e5e7981a1069 // indirect
)

replace github.com/rsteube/carapace => ../carapace
replace github.com/rsteube/invoke-completion => ../invoke-completion
