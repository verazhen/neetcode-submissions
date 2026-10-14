/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    maxDiameter := 0

    var getHeight func(*TreeNode) int
    getHeight = func(node *TreeNode) int {
        if node == nil {
            return 0
        }

        leftHeight := getHeight(node.Left)
        rightHeight := getHeight(node.Right)

        // Update the longest path found so far
        maxDiameter = max(maxDiameter, leftHeight+rightHeight)

        // Return the height of the current subtree
        return 1 + max(leftHeight, rightHeight)
    }

    getHeight(root)
    return maxDiameter
}
