/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    return validateBST(root, -1000000001, 1000000001)
}

func validateBST(node *TreeNode, minVal, maxVal int) bool {
    if node == nil {
        return true
    }

    if node.Val <= minVal || node.Val >= maxVal {
        return false
    }

    return validateBST(node.Left, minVal, node.Val) &&
        validateBST(node.Right, node.Val, maxVal)
}
