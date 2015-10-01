package main

import (
	"bufio"
	"os"

	"github.com/calebthompson/ftree/tree"
)

func main() {
	r := bufio.NewScanner(os.Stdin)
	lines := []string{}
	for r.Scan() {
		lines = append(lines, r.Text())
	}

	t := tree.Tree(lines, "/")
	t.Print(0, nil)
}
