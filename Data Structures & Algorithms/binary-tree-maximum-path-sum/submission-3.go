/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
	//get two sume from both line
	//add node value
	//add up both line and update the max
    //return bigger sub-line sum only

	maxSum := root.Val
	var updateMax func(root *TreeNode) int
	updateMax = func(root *TreeNode) int{
		if root == nil{
			return 0
		}

		leftSum := updateMax(root.Left)
		rightSum := updateMax(root.Right)

		bigger := root.Val
		if !(leftSum<=0 && rightSum<=0){
			bigger += max(leftSum,rightSum)
		}

		maxSum = max(maxSum,leftSum+rightSum+ root.Val)
		maxSum = max(maxSum,bigger)

		return  bigger
	}
	updateMax(root)
	return maxSum
}
