package models

import c "aoc/common"

type Node struct {
	parent      *Node
	left, right *Node
	value       int64
	size        int
	height      int
}

func (n *Node) GetValue() int64 {
	return n.value
}

func (n *Node) rebalance() *Node {
	is_unbalanced := true
	node := n
	for is_unbalanced {
		node.update()
		b := get_balance(node)
		is_unbalanced = b < -1 || 1 < b
		if !is_unbalanced {
			continue
		}

		// left subtree is higher
		if b < -1 {
			if get_balance(node.left) > 0 {
				node.left = node.left.swap(node.left.right)
			}
			node = node.swap(node.left)
		} else /* right subtree is higher */ {
			if get_balance(node.right) < 0 {
				node.right = node.right.swap(node.right.left)
			}
			node = node.swap(node.right)
		}
	}
	return node
}

// Gets the index of this node w.r.t. the subtree rooted at this node
func (node *Node) get_my_relative_index() int {
	return get_size(node.left)
}

func (node *Node) set_left(child *Node) {
	node.left = child
	set_parent(child, node)
}

func (node *Node) set_right(child *Node) {
	node.right = child
	set_parent(child, node)
}

type find_stack_instance struct {
	parent_node *Node
	next_ptr    **Node
	next_node   *Node
}

func rebuild_stack(stack *c.Stack[find_stack_instance]) {
	for !stack.IsEmpty() {
		fsi, _ := stack.Pop()
		next := fsi.next_node
		if next != nil {
			next = next.rebalance()
			next.parent = fsi.parent_node
		}
		*fsi.next_ptr = next
	}
}

func find(fsi find_stack_instance, index int) *c.Stack[find_stack_instance] {
	node := fsi.next_node
	stack := c.MakeStack[find_stack_instance](fsi)
	for node != nil && node.get_my_relative_index() != index {
		var next_ptr **Node = nil
		if node.get_my_relative_index() < index {
			index = index - (node.get_my_relative_index() + 1)
			next_ptr = &node.right
		} else {
			next_ptr = &node.left
		}
		stack.Push(find_stack_instance{parent_node: node, next_ptr: next_ptr, next_node: *next_ptr})
		node = *next_ptr
	}
	return stack
}

func (n *Node) Clear() {
	n.left, n.right, n.parent = nil, nil, nil
	n.update()
}

func MakeNode(val int64) *Node {
	return &Node{
		value: val,
	}
}

// Gets the index of this node w.r.t. the whole AVL tree
func (node *Node) GetMyIndex() int {
	position := get_size(node.left)
	current := node
	for current.parent != nil {
		if current.parent.right == current {
			position += 1 + get_size(current.parent.left)
		}
		current = current.parent
	}
	return position
}

func in_order(node *Node, f func(*Node)) {
	if node == nil {
		return
	} else {
		in_order(node.left, f)
		f(node)
		in_order(node.right, f)
	}
}

func set_parent(node *Node, new_parent *Node) {
	if node != nil {
		node.parent = new_parent
	}
}

// Update height & size based on the children nodes
func (node *Node) update() *Node {
	node.height = 1 + c.Max(get_height(node.left), get_height(node.right))
	node.size = 1 + get_size(node.left) + get_size(node.right)
	return node
}

// Swap parent with its child
func (current *Node) swap(child *Node) *Node {
	if current.left == child {
		child.right, current.left = current, child.right
		set_parent(child.right, child)
		set_parent(current.left, current)
	} else if current.right == child {
		child.left, current.right = current, child.left
		set_parent(child.left, child)
		set_parent(current.right, current)
	} else {
		return current
	}
	current.parent, child.parent = child, current.parent
	current.update()
	return child.update()
}

func get_balance(node *Node) int {
	if node == nil {
		return 0
	} else {
		return get_height(node.right) - get_height(node.left)
	}
}

func get_size(node *Node) int {
	if node == nil {
		return 0
	} else {
		return node.size
	}
}

func get_height(node *Node) int {
	if node == nil {
		return 0
	} else {
		return node.height
	}
}
