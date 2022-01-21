function levelOrder(root: TreeNode | null): number[][] {
    if (root == null){
        return []
    }
    let array: number[][] = []
    let row: TreeNode[] = [root]
    let nodeArray: TreeNode[]
    let valueArray: number[]
    let node: TreeNode
    while (row.length != 0) {
        nodeArray = []
        valueArray = []
        for(node of row){
            valueArray.push(node.val)
            if (node.left != null){
                nodeArray.push(node.left)
            }
            if (node.right != null){
                nodeArray.push(node.right)
            }
        }
        array.push(valueArray)
        row = nodeArray
    }
    return array
};