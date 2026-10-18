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

    var inorder func(*TreeNode)
    inorder = func(node *TreeNode) {
        if node == nil || count == k {
            return
        }

        inorder(node.Left)
        if count == k {
            return
        }

        count++
        if count == k {
            answer = node.Val
            return
        }

        inorder(node.Right)
    }

    inorder(root)
    return answer
}