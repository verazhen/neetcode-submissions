/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {

	if isSameTree(root,subRoot){
		return true
	}

	if root == nil{
		return false
	}

	return isSubtree(root.Left,subRoot) || isSubtree(root.Right,subRoot)

}

func isSameTree(r *TreeNode, s *TreeNode) bool {
	if r == nil && s == nil{
		return true
	}else if r == nil || s == nil{
		return false
	}else if r.Val != s.Val{
		return false
	}

	return isSameTree(r.Left,s.Left) && isSameTree(r.Right,s.Right)
}

