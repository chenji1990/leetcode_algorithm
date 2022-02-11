function lowestCommonAncestor(root: TreeNode | null, p: TreeNode | null, q: TreeNode | null): TreeNode | null {
    if (root == null || p == null || q == null){
        return root
    }
    if (p.val > q.val){
        let temp = p
        p = q
        q = temp
    }
    while(root != null){
        if (root.val > q.val){
            root = root.left
            continue
        }
        if (root.val < p.val){
            root = root.right
            continue
        }
        return root
    }
    return root
};