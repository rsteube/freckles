package main

import "github.com/carapace-sh/freckles/cmd/freckles/cmd"

var version = "develop"

func main() {
	cmd.Execute(version)
}
