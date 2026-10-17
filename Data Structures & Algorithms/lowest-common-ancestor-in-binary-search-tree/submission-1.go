/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
    small := 0
	big := 0
	if p.Val < q.Val{
		small = p.Val
		big = q.Val
	}else{
		small = q.Val
		big = p.Val
	}
	
	for root != nil{
		if small <= root.Val && big >= root.Val{
			return root
		}else if small < root.Val && big < root.Val{
			root = root.Left
		}else if small > root.Val && big > root.Val{
			root = root.Right
		}
	}
	return nil
}
