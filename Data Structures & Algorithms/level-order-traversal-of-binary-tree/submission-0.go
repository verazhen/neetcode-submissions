/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {

	result := [][]int{}
    var pushNode func (root *TreeNode, levelIndex int)
	pushNode = func (root *TreeNode, levelIndex int){
		if root == nil{
			return
		}
		if len(result) <= levelIndex{
			result = append(result,[]int{})
		}
		result[levelIndex] = append(result[levelIndex],root.Val)
		pushNode(root.Left,levelIndex+1)
		pushNode(root.Right,levelIndex+1)
		return
	}
	pushNode(root,0)
	return result
}


