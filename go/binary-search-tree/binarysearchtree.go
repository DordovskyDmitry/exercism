package binarysearchtree

type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

func Bst(n int) *SearchTreeData {
	return &SearchTreeData{data: n}
}

func (bst *SearchTreeData) Insert(n int) {
	if bst.data >= n {
		if bst.left != nil {
			bst.left.Insert(n)
		} else {
			bst.left = Bst(n)
		}
	} else {
		if bst.right != nil {
			bst.right.Insert(n)
		} else {
			bst.right = Bst(n)
		}
	}
}

func (bst *SearchTreeData) MapString(f func(int) string) []string {
	res := make([]string, 0)
	if bst.left != nil {
		res = append(bst.left.MapString(f), res...)
	}
	res = append(res, f(bst.data))
	if bst.right != nil {
		res = append(res, bst.right.MapString(f)...)
	}
	return res
}

func (bst *SearchTreeData) MapInt(f func(int) int) []int {
	res := make([]int, 0)
	if bst.left != nil {
		res = append(bst.left.MapInt(f), res...)
	}
	res = append(res, f(bst.data))
	if bst.right != nil {
		res = append(res, bst.right.MapInt(f)...)
	}
	return res
}
