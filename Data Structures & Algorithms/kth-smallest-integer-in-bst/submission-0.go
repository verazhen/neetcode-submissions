/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    count := 0
	answer := -1
	var inorder func(root *TreeNode) 
	inorder = func(root *TreeNode) {
		if root == nil{
			return
		}

		inorder(root.Left)
		count++
		if count == k {
			answer = root.Val
		} 
		inorder(root.Right)
	}
	inorder(root)
	return answer

}

