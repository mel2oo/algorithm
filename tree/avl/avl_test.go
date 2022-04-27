package avl

import (
	"testing"
)

type Int int

func (t Int) Equals(k Key) bool {
	return t == k
}

func (t Int) Less(k Key) bool {
	return t < k.(Int)
}

func TestAVL(t *testing.T) {
	tree := &AVLTree{}

	tree.Insert(Int(1), 1)
	tree.Insert(Int(2), 2)
	tree.Insert(Int(3), 3)
	tree.Insert(Int(4), 4)
	tree.Insert(Int(5), 5)
	tree.Insert(Int(6), 6)

	tree.DisplayInOrder()
}
