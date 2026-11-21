// Package ds holds the shared LeetCode data structures and test helpers.
// In a solution, use type aliases so signatures match LeetCode:
//
//	type ListNode = ds.ListNode
//	type TreeNode = ds.TreeNode
package ds

// ListNode is LeetCode's singly linked list node.
type ListNode struct {
	Val  int
	Next *ListNode
}

// TreeNode is LeetCode's binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewList builds a linked list from values: NewList(1,2,3) → 1→2→3.
func NewList(vals ...int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

// ListToSlice converts a linked list back to a slice (nil list → empty slice).
func ListToSlice(head *ListNode) []int {
	out := []int{}
	for ; head != nil; head = head.Next {
		out = append(out, head.Val)
	}
	return out
}

// NewTree builds a tree from LeetCode level-order notation, with nil for missing nodes.
// Example: NewTree(3, 9, 20, nil, nil, 15, 7)
func NewTree(vals ...any) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]
		if i < len(vals) && vals[i] != nil {
			node.Left = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(vals) && vals[i] != nil {
			node.Right = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

// TreeToSlice converts a tree to LeetCode level-order notation (trailing nils trimmed).
func TreeToSlice(root *TreeNode) []any {
	out := []any{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			out = append(out, nil)
			continue
		}
		out = append(out, node.Val)
		queue = append(queue, node.Left, node.Right)
	}
	for len(out) > 0 && out[len(out)-1] == nil {
		out = out[:len(out)-1]
	}
	return out
}

// FindNode returns the first node with value v (handy for LCA-style tests).
func FindNode(root *TreeNode, v int) *TreeNode {
	if root == nil || root.Val == v {
		return root
	}
	if n := FindNode(root.Left, v); n != nil {
		return n
	}
	return FindNode(root.Right, v)
}
