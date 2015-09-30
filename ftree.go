package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	r := bufio.NewScanner(os.Stdin)
	lines := []string{}
	for r.Scan() {
		lines = append(lines, r.Text())
	}

	var parts []node
	for _, line := range lines {
		parts = append(parts, buildTree(line))
	}

	var tree node
	for _, part := range parts {
		tree = merge(part, tree)
	}

	tree.Print(0)
}

func buildTree(line string) node {
	var subtree = node{}
	split := strings.SplitN(line, "/", 2)
	if len(split) > 1 {
		subtree[split[0]] = buildTree(split[1])
	} else {
		subtree[split[0]] = node{}
	}
	return subtree
}

func merge(src node, dest node) node {
	if dest == nil {
		dest = node{}
	}
	for k, v := range src {
		dest[k] = merge(v, dest[k])
	}
	return dest
}

type node map[string]node

func (n node) Print(indent int) {
	var keys []string
	for k := range n {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, part := range keys {
		for i := 0; i < indent; i++ {
			fmt.Print("│   ")
		}
		fmt.Printf("├── %s\n", part)
		n[part].Print(indent + 1)
	}
}
