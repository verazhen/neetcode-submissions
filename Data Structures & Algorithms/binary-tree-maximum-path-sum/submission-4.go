func maxPathSum(root *TreeNode) int {
    maxSum := root.Val

    var dfs func(node *TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }

        leftGain := max(0, dfs(node.Left))
        rightGain := max(0, dfs(node.Right))

        // 最大路徑可以同時包含左右子樹
        pathSum := node.Val + leftGain + rightGain
        maxSum = max(maxSum, pathSum)

        // 回傳給父節點時，只能選擇其中一側
        return node.Val + max(leftGain, rightGain)
    }

    dfs(root)
    return maxSum
}