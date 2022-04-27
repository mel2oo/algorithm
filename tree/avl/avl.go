package avl

import "fmt"

// AVLNode Key Interface
type Key interface {
	Equals(Key) bool
	Less(Key) bool
}

// AVLTree structure
type AVLTree struct {
	root *AVLNode
}

func (t *AVLTree) Insert(key Key, value interface{}) {
	t.root = t.root.add(key, value)
}

func (t *AVLTree) Remove(key Key) {
	t.root = t.root.remove(key)
}

func (t *AVLTree) Update(oldKey Key, newKey Key, newValue int) {
	t.root = t.root.remove(oldKey)
	t.root = t.root.add(newKey, newValue)
}

func (t *AVLTree) Search(key Key) (node *AVLNode) {
	return t.root.search(key)
}

func (t *AVLTree) DisplayInOrder() {
	t.root.displayNodesInOrder()
}

// AVLNode structure
type AVLNode struct {
	key   Key
	Value interface{}

	height int
	left   *AVLNode
	right  *AVLNode
}

func (n *AVLNode) add(key Key, value interface{}) *AVLNode {
	if n == nil {
		return &AVLNode{key, value, 1, nil, nil}
	}

	if key.Equals(n.key) {
		n.Value = value
	} else if key.Less(n.key) {
		n.left = n.left.add(key, value)
	} else {
		n.right = n.right.add(key, value)
	}
	return n.rebalanceTree()
}

func (n *AVLNode) remove(key Key) *AVLNode {
	if n == nil {
		return nil
	}
	if key.Equals(n.key) {
		if n.left != nil && n.right != nil {
			rightMinNode := n.right.findSmallest()
			n.key = rightMinNode.key
			n.Value = rightMinNode.Value
			n.right = n.right.remove(rightMinNode.key)
		} else if n.left != nil {
			n = n.left
		} else if n.right != nil {
			n = n.right
		} else {
			n = nil
			return n
		}
	} else if key.Less(n.key) {
		n.left = n.left.remove(key)
	} else {
		n.right = n.right.remove(key)
	}

	return n.rebalanceTree()
}

func (n *AVLNode) search(key Key) *AVLNode {
	if n == nil {
		return nil
	}

	if key.Equals(n.key) {
		return n
	} else if key.Less(n.key) {
		return n.left.search(key)
	} else {
		return n.right.search(key)
	}
}

func (n *AVLNode) displayNodesInOrder() {
	if n.left != nil {
		n.left.displayNodesInOrder()
	}
	fmt.Print(n.key, " ")
	if n.right != nil {
		n.right.displayNodesInOrder()
	}
}

func (n *AVLNode) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *AVLNode) recalculateHeight() {
	n.height = 1 + max(n.left.getHeight(), n.right.getHeight())
}

func (n *AVLNode) rebalanceTree() *AVLNode {
	if n == nil {
		return n
	}
	n.recalculateHeight()

	balanceFactor := n.left.getHeight() - n.right.getHeight()
	if balanceFactor == -2 {
		if n.right.left.getHeight() > n.right.right.getHeight() {
			n.right = n.right.rotateRight()
		}
		return n.rotateLeft()
	} else if balanceFactor == 2 {
		if n.left.right.getHeight() > n.left.left.getHeight() {
			n.left = n.left.rotateLeft()
		}
		return n.rotateRight()
	}
	return n
}

func (n *AVLNode) rotateLeft() *AVLNode {
	newRoot := n.right
	n.right = newRoot.left
	newRoot.left = n

	n.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

func (n *AVLNode) rotateRight() *AVLNode {
	newRoot := n.left
	n.left = newRoot.right
	newRoot.right = n

	n.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

func (n *AVLNode) findSmallest() *AVLNode {
	if n.left != nil {
		return n.left.findSmallest()
	} else {
		return n
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
