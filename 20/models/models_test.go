package models

import (
	c "aoc/common"
	ts "aoc/testers"
	"testing"
)

func TestD20_AVL_tree(t *testing.T) {
	avl_tree := MakeAVL_Tree()
	assertor := AVL_tree_assertion_wrapper(t, avl_tree)
	for i := range c.RangeInclusive(0, 5) {
		avl_tree.InsertAtIndex(i, MakeNode(int64(i)))
		ts.Assert(t, is_properly_linked(avl_tree.tree))
	}
	// 0,1,2,3,4,5
	avl_tree.InsertAtIndex(2, MakeNode(10))
	assertor([]int64{0, 1, 10, 2, 3, 4, 5})

	avl_tree.InsertAtIndex(2, MakeNode(20))
	assertor([]int64{0, 1, 20, 10, 2, 3, 4, 5})

	avl_tree.InsertAtIndex(2, MakeNode(30))
	assertor([]int64{0, 1, 30, 20, 10, 2, 3, 4, 5})

	avl_tree.InsertAtIndex(0, MakeNode(-101))
	assertor([]int64{-101, 0, 1, 30, 20, 10, 2, 3, 4, 5})

	avl_tree.InsertAtIndex(7, MakeNode(101))
	assertor([]int64{-101, 0, 1, 30, 20, 10, 2, 101, 3, 4, 5})

	avl_tree.RemoveIndex(2)
	assertor([]int64{-101, 0, 30, 20, 10, 2, 101, 3, 4, 5})

	avl_tree.RemoveIndex(0)
	assertor([]int64{0, 30, 20, 10, 2, 101, 3, 4, 5})

	avl_tree.RemoveIndex(7)
	assertor([]int64{0, 30, 20, 10, 2, 101, 3, 5})
}

func AVL_tree_assertion_wrapper(t *testing.T, avl *AVL_Tree) func([]int64) {
	return func(expected []int64) {
		ts.AssertEqualWithEqFunc(t, avl.GetAsArray(), expected, c.ArrayEqual[int64])
		ts.Assert(t, is_properly_linked(avl.tree))
		underlying_nodes := make([]*Node, 0)
		in_order(avl.tree, func(n *Node) { underlying_nodes = append(underlying_nodes, n) })

		for i := range expected {
			ts.AssertEqual(t, underlying_nodes[i].GetMyIndex(), i)
		}
	}
}

func is_properly_linked(node *Node) bool {
	if node == nil {
		return true
	}
	return (node.left == nil || node.left.parent == node) && (node.right == nil || node.right.parent == node) && is_properly_linked(node.left) && is_properly_linked(node.right)
}
