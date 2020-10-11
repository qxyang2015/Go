package MirrorTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	MirrorTree(root.Left)
	MirrorTree(root.Right)
	return root
}
