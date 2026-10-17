/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	queue := []*TreeNode{root}
	result := [][]int{}
	
	for root !=nil && len(queue) > 0 {
    	levelSize := len(queue)
		sub := []int{}
		for levelSize > 0{
			q := queue[0]
			sub = append(sub,q.Val)

			queue = queue[1:]
			if q.Left != nil{
				queue = append(queue,q.Left)
			}
			if q.Right != nil{
				queue = append(queue,q.Right)
			}
			levelSize--
		}
		result = append(result,sub)
	}

	return result
}
