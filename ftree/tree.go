package ftree

import "strings"

// Tree splits each line in lines on delim and builds a Node: a tree data
// structure of maps with string keys.
func Tree(lines []string, delim string) Node {
	var parts []Node
	for _, line := range lines {
		parts = append(parts, buildTree(line, delim))
	}

	var tree Node
	for _, part := range parts {
		tree = merge(part, tree)
	}

	return tree
}

func buildTree(line, delim string) Node {
	var subtree = Node{}
	split := strings.SplitN(line, delim, 2)
	if len(split) > 1 {
		subtree[split[0]] = buildTree(split[1], delim)
	} else {
		subtree[split[0]] = Node{}
	}
	return subtree
}

func merge(src Node, dest Node) Node {
	if dest == nil {
		dest = Node{}
	}
	for k, v := range src {
		dest[k] = merge(v, dest[k])
	}
	return dest
}
