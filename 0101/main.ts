function isSymmetric(root: TreeNode | null): boolean {
    return check(root, root)
};

function check(p: TreeNode | null, q: TreeNode | null): boolean{
    if (p == null && q == null){
        return true
    }
    if (p == null || q == null){
        return false
    }
    return p.val == q.val && check(p.left, q.right) && check(p.right, q.left)
}