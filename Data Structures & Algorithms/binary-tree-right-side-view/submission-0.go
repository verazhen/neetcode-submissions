/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	if root == nil{
		return []int{}
	}
	
	queue := []*TreeNode{root}
	result := []int{}
	for len(queue) > 0{
		queueSize := len(queue)
		for queueSize > 0{
			q := queue[0]
			if queueSize == 1{
				result = append(result,q.Val)
			}
			queue = queue[1:]
			if q.Left !=nil{
				queue = append(queue, q.Left)
			}
			if q.Right !=nil{
				queue = append(queue, q.Right)
			}
			queueSize--
		}
	}
	return result
}
