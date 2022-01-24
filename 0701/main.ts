function insertIntoBST(root: TreeNode | null, val: number): TreeNode | null {
    if (root == null){
        return new TreeNode(val)
    }
    if (root.val < val){
        root.right = insertIntoBST(root.right, val)
    } else if (root.val > val) {
        root.left = insertIntoBST(root.left, val)
    }
    return root
};