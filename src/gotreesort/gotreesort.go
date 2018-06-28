package gotreesort

import (
	"fmt"
)

type tree struct {
	value int
	left, right *tree
}
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = Add(root, v)
	}

	AppendValues(values[:0], root)
	fmt.Println("---> ", values)
}

func AppendValues(values []int, t *tree) []int {
	if t != nil {
		values = AppendValues(values, t.left)
		values = append(values, t.value)
		values = AppendValues(values, t.right)
	}
	return values
}

func Add(t *tree, value int) *tree {
	if t == nil {
		t =new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = Add(t.left,value)
	} else {
		t.right = Add(t.right, value)
	}
	return t
}