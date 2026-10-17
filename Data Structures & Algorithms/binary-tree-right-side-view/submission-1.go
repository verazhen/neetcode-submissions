/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
    result := []int{}
    var pushNode func (root *TreeNode, depth int)
	pushNode = func (root *TreeNode, depth int){
		if root == nil{
			return
		}
		if len(result) <= depth{
			result = append(result,root.Val)
		}
		pushNode(root.Right,depth+1)
		pushNode(root.Left,depth+1)
		return
	}
	pushNode(root,0)
	return result
}
