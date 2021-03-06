package tree

import (
	"fmt"
	"sort"
)

// Node is a tree of string. Leaves are empty Nodes or nil.
type Node map[string]Node

// Print recurses through n and its member Nodes to pretty-print a Node as a
// tree such as:
//
//     ├── ftree
//     │   ├── node.go
//     │   ├── tree.go
//     │   └── tree_test.go
//     └── main.go
//
// Print should initially be called as n.Print(0, nil). The arguments are
// primarily interesting during recursion.
func (n Node) Print(indent int, leaders []string) {
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
