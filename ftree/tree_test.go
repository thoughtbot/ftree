package ftree

import (
	"reflect"
	"testing"
)

func TestTree(t *testing.T) {
	lines := []string{
		"caleb/code/ftree",
		"caleb/books",
	}

	want := Node{
		"caleb": Node{
			"code":  Node{"ftree": Node{}},
			"books": Node{},
		},
	}

	got := Tree(lines, "/")

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Expected \n%#v, got \n%#v", want, got)
	}
}
