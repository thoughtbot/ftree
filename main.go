package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/calebthompson/ftree/tree"
)

const (
	version = "v0.2.0"
	help    = `Pipe output from some command that lists a single file per line to ftree.
	find . | ftree
	git ls-files {app,spec}/models | ftree
	lsrc | cut -d : -f 1 | ftree
`
)

func main() {
	handleArguments()

	r := bufio.NewScanner(os.Stdin)
	lines := []string{}
	for r.Scan() {
		lines = append(lines, r.Text())
	}

	t := tree.Tree(lines, "/")
	t.Print(0, nil)
}

func handleArguments() {
	for _, a := range os.Args[1:] {
		switch a {
		case "--version", "-v":
			fmt.Printf("ftree %s\n", version)
			os.Exit(0)
		case "--help", "-h":
			fmt.Print(help)
			os.Exit(0)
		default:
			fmt.Fprintf(os.Stderr, "No such argument: %v\n", a)
			os.Exit(1)
		}
	}
}
