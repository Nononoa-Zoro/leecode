package solutions

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// lowestCommonAncestor 二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	//  左子树没有p和q
	if left == nil {
		return right
	}
	// 右子树没有p和q
	if right == nil {
		return left
	}
	// 左右都有，说明p和q在两侧，此时root节点就是p和q的根节点
	return root
}
