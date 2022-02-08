function invertTree(root: TreeNode | null): TreeNode | null {
    if (root == null) {
        return null
    }
    let temp = invertTree(root.left)
    root.left = invertTree(root.right)
    root.right = temp
    return root
};