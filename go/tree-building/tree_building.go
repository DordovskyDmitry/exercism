package tree

import (
	"sort"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

type Mismatch struct{}

func (m Mismatch) Error() string {
	return "c"
}

func Build(records []Record) (*Node, error) {
	length := len(records)
	if length == 0 {
		return nil, nil
	}
	// The idea was peeped rickh94
	sort.SliceStable(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	dependencies := make([][]int, length)
	for _, record := range records {
		id := record.ID
		parent := record.Parent
		if id >= length || parent >= length {
			return nil, Mismatch{}
		}
		if id == 0 && parent == 0 {
			continue
		}
		if parent >= id && parent != 0 {
			return nil, Mismatch{}
		}
		dependencies[parent] = append(dependencies[parent], id)
	}

	root := &Node{ID: 0}
	buildSubtree(root, dependencies)
	return root, nil
}

func buildSubtree(root *Node, dependencies [][]int) {
	for _, child := range dependencies[root.ID] {
		childNode := &Node{ID: child}
		root.Children = append(root.Children, childNode)
		buildSubtree(childNode, dependencies)
	}
}
