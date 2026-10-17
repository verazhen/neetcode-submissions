/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    if root == nil{
		return 0
	}
	
	var compareWithMax func(root *TreeNode, maxVal int) int
	compareWithMax = func(root *TreeNode, maxVal int) int{
		count := 0
		if root== nil{
			return 0
		}

		if root.Val >= maxVal{
			count ++ 
			maxVal = root.Val
		}
		count += compareWithMax(root.Left,maxVal)
		count += compareWithMax(root.Right,maxVal)
		return count
	}
	
	return compareWithMax(root,root.Val)
}
