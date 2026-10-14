/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    dia,maximum :=diameter(root)
	return max(dia-1,maximum)
}

func diameter(root *TreeNode)(int,int){
	if root == nil{
		return 0,0
	}

	left, leftMax := diameter(root.Left)
	right, rightMax := diameter(root.Right)
	return max(left,right)+1,max(max(leftMax,rightMax),left+right)
}
