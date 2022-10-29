package deletenodeinabst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// recursion1
func deleteNode1(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val == key {
		if root.Left == nil {
			return root.Right
		}

		if root.Right == nil {
			return root.Left
		}

		min := getMin(root.Right)
		root.Val = min.Val
		root.Right = deleteNode1(root.Right, min.Val)
	} else if root.Val > key {
		root.Left = deleteNode1(root.Left, key)
	} else {
		root.Right = deleteNode1(root.Right, key)
	}

	return root
}

func getMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	for root.Left != nil {
		root = root.Left
	}

	return root
}

// recursion2
func deleteNode2(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val == key {
		if root.Left == nil {
			return root.Right
		}

		if root.Right == nil {
			return root.Left
		}

		cur := root.Right
		for cur.Left != nil {
			cur = cur.Left
		}

		cur.Left = root.Left
		return root.Right
	} else if root.Val > key {
		root.Left = deleteNode2(root.Left, key)
	} else {
		root.Right = deleteNode2(root.Right, key)
	}

	return root
}
