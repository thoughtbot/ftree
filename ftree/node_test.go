package ftree

func ExamplePrint() {
	tree := Node{
		"caleb": Node{
			"code": Node{"ftree": Node{}},
			"books": Node{
				"Learning Go":        Node{},
				"The Little Go Book": Node{},
			},
		},
	}

	tree.Print(0, nil)

	// Output:
	// └── caleb
	//     ├── books
	//     │   ├── Learning Go
	//     │   └── The Little Go Book
	//     └── code/ftree
}
