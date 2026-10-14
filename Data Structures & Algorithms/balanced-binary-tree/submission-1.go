/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    if root == nil{
		return true
	}
	
	return !(getHeight(root)==-1)
}
func getHeight(root *TreeNode) int {
	if root == nil{
		return 0
	}

	left := getHeight(root.Left)
	right := getHeight(root.Right)

	deduct := left - right
	if deduct < -1 || deduct > 1 || left == -1 || right == -1{
		return -1
	} 

	return max(left,right)+1
}
