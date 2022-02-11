function postorderTraversal(root: TreeNode | null): number[] {
    let array: number[] = []
    function order(node: TreeNode | null){
        if (node == null){
            return
        }
        order(node.left)
        order(node.right)
        array.push(node.val)
    }
    order(root)
    return array
};