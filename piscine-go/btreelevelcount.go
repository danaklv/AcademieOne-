package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	LeftSide := BTreeLevelCount(root.Left)
	RightSide := BTreeLevelCount(root.Right)
	if LeftSide > RightSide {
		return LeftSide + 1
	}
	return RightSide + 1
}
