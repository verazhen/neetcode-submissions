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
	
	count := 0
	var compareWithMax func(root *TreeNode, maxVal int)
	compareWithMax = func(root *TreeNode, maxVal int){
		if root== nil{
			return
		}

		if root.Val >= maxVal{
			count ++ 
			maxVal = root.Val
		}
		compareWithMax(root.Left,maxVal)
		compareWithMax(root.Right,maxVal)
	}
	compareWithMax(root,root.Val)
	return count
}
