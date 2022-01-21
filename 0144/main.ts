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