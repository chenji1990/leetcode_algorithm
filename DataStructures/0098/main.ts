function isValidBST(root: TreeNode | null): boolean {
    return isValid(root, -1e32, 1e32)
};

function isValid(node: TreeNode | null, lower: number, upper: number): boolean {
    if (node == null){
        return true
    }
    if (node.val <= lower || node.val >= upper){
        return false
    }
    return isValid(node.left, lower, node.val) && isValid(node.right, node.val, upper)
};