/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    return isValidFunc(root,-1000000001,1000000001)
}
func isValidFunc(root *TreeNode, minimum, maximum int)bool{
	if root == nil{
		return true
	}else if root.Val <= minimum || root.Val >= maximum{
		return false
	}
	
	leftMin := minimum
	leftMax := root.Val
	rightMin := root.Val
	rightMax := maximum
	
	return isValidFunc(root.Left,leftMin,leftMax) && isValidFunc(root.Right,rightMin,rightMax) 
}
