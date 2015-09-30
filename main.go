package main

import (
	"bufio"
	"os"

	"github.com/calebthompson/ftree/ftree"
)

func main() {
	r := bufio.NewScanner(os.Stdin)
	lines := []string{}
	for r.Scan() {
		lines = append(lines, r.Text())
	}

	t := ftree.Tree(lines, "/")
	t.Print(0, nil)
}
