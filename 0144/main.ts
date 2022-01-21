
function preorderTraversal(root: TreeNode | null): number[] {
    let array: number[] = []
    function order(node: TreeNode | null){
        if (node == null){
            return
        }
        array.push(node.val)
        order(node.left)
        order(node.right)
    }
    order(root)
    return array
};