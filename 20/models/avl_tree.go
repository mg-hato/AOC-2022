package models

type AVL_Tree struct {
	tree *Node
}

func MakeAVL_Tree() *AVL_Tree {
	return &AVL_Tree{}
}

func (avl AVL_Tree) Size() int {
	return get_size(avl.tree)
}

func (avl *AVL_Tree) InsertAtIndex(index int, node *Node) {
	node.Clear()
	stack := find(find_stack_instance{next_ptr: &avl.tree, next_node: avl.tree}, index)

	fsi, _ := stack.Pop()
	stack.Push(find_stack_instance{parent_node: fsi.parent_node, next_ptr: fsi.next_ptr, next_node: node})
	stack.Push(find_stack_instance{parent_node: node, next_ptr: &node.right, next_node: fsi.next_node})

	if fsi.next_node != nil {
		stack.Push(find_stack_instance{parent_node: node, next_ptr: &node.left, next_node: fsi.next_node.left})
		fsi.next_node.left = nil
	}

	rebuild_stack(stack)
}

func (avl *AVL_Tree) RemoveIndex(index int) {
	stack := find(find_stack_instance{next_ptr: &avl.tree, next_node: avl.tree}, index)

	// If node not found: nothing to do
	if top, _ := stack.Top(); top.next_node == nil {
		return
	}

	// If node to be removed has at least one nil-child: successor is easy to be determined
	if top, _ := stack.Top(); top.next_node.left == nil || top.next_node.right == nil {
		fsi, _ := stack.Pop()
		successor := fsi.next_node.left
		if successor == nil {
			successor = fsi.next_node.right
		}
		fsi.next_node = successor
		stack.Push(fsi)
		rebuild_stack(stack)
		return
	}

	// Otherwise, successor is "lowest" node in the right subtree of the node to be removed
	fsi, _ := stack.Pop()

	var succ_right *Node = nil
	succ_stack := find(find_stack_instance{next_ptr: &succ_right, next_node: fsi.next_node.right}, 0)
	succ_fsi, _ := succ_stack.Pop()
	successor := succ_fsi.next_node
	succ_fsi.next_node = successor.right
	succ_stack.Push(succ_fsi)
	rebuild_stack(succ_stack)

	successor.set_right(succ_right)
	successor.set_left(fsi.next_node.left)

	fsi.next_node = successor
	stack.Push(fsi)
	rebuild_stack(stack)
}

func (avl *AVL_Tree) GetAsArray() []int64 {
	arr := make([]int64, 0)
	in_order(avl.tree, func(n *Node) { arr = append(arr, n.value) })
	return arr
}
