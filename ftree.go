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

	tree.Print(0, nil)
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

func (n node) Print(indent int, leaders []string) {
	var keys []string
	for k := range n {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i, part := range keys {
		if part == "" {
			n[part].Print(indent, leaders)
			continue
		}

		for _, l := range leaders {
			fmt.Print(l)
		}

		// For chains of single element maps, such as
		// { "Users": { "caleb": { "code": {}, "books": {} } },
		// combine to { "Users/caleb": { "code": {}, "books": {} } }.
		for {
			if len(n[part]) != 1 {
				break
			}
			for k, v := range n[part] {
				part += "/" + k
				n[part] = v
			}
		}

		if i == len(keys)-1 {
			fmt.Printf("%s %s\n", "└──", part)
			n[part].Print(indent+1, append(leaders, "    "))
		} else {
			fmt.Printf("%s %s\n", "├──", part)
			n[part].Print(indent+1, append(leaders, "│   "))
		}
	}
}
