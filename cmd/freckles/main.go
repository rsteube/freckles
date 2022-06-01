package main

import "github.com/rsteube/freckles/cmd/freckles/cmd"

var version = "develop"

func main() {
	cmd.Execute(version)
}
