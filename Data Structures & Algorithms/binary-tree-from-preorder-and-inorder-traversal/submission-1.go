func buildTree(preorder []int, inorder []int) *TreeNode {
    // Map each node value to its index in inorder.
    inorderIndex := make(map[int]int, len(inorder))
    for i, val := range inorder {
        inorderIndex[val] = i
    }

    // Track the next root to build in preorder.
    preorderIndex := 0

    var buildSubtree func(left, right int) *TreeNode

    buildSubtree = func(left, right int) *TreeNode {
        if left > right {
            return nil
        }

        // Create the root of the current subtree.
        rootVal := preorder[preorderIndex]
        preorderIndex++

        rootIndex := inorderIndex[rootVal]

        return &TreeNode{
			Val: rootVal,
			Left: buildSubtree(left, rootIndex-1),
			Right: buildSubtree(rootIndex+1, right),
		}
    }

    return buildSubtree(0, len(inorder)-1)
}