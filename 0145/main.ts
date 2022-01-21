class TreeNode {
    val: number
    left: TreeNode | null
    right: TreeNode | null
    constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
        this.val = (val === undefined ? 0 : val)
        this.left = (left === undefined ? null : left)
        this.right = (right === undefined ? null : right)
    }
}

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