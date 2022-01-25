public class TreeNode {
    public var val: Int
    public var left: TreeNode?
    public var right: TreeNode?
    public init() { self.val = 0; self.left = nil; self.right = nil; }
    public init(_ val: Int) { self.val = val; self.left = nil; self.right = nil; }
    public init(_ val: Int, _ left: TreeNode?, _ right: TreeNode?) {
        self.val = val
        self.left = left
        self.right = right
    }
}

class Solution {
    func findTarget(_ root: TreeNode?, _ k: Int) -> Bool {
        var hashmap = [Int: Bool]()
        func find(_ node: TreeNode?, _ target: Int) -> Bool{
            guard let node = node else {
                return false
            }
            if let _ = hashmap[target - node.val]{
                return true
            }
            hashmap[node.val] = true
            return find(node.left, target) || find(node.right, target)
        }
        return find(root, k)
    }
}