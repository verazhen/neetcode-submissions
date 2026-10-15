/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
	index := 0
    var buildFunc func(inorder []int)*TreeNode
	buildFunc = func(inorder []int)*TreeNode{
		if index >= len(preorder){
			return nil
		}else if len(inorder) == 0{
			return nil
		}
		root := preorder[index]
		left := []int{}
		right := []int{}
		for i,node:=range inorder{
			if node == root{
				left = inorder[:i]
				right = inorder[i+1:]
				index++
				break
			}

		}

		return &TreeNode{
			Val: root,
			Left: buildFunc(left),
			Right: buildFunc(right),
		}
	}
	return buildFunc(inorder)
}

