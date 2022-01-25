function findTarget(root: TreeNode | null, k: number): boolean {
    let hashmap: Map<number, boolean> = new Map()
    function find(node: TreeNode | null, target: number): boolean {
        if (node == null){
            return false
        }
        if (hashmap.get(target - node.val)){
            return true
        }
        hashmap.set(node.val, true)
        return find(node.left, target) || find(node.right, target)
    };
    return find(root, k)
};