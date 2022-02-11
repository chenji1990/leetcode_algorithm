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
    func isValidBST(_ root: TreeNode?) -> Bool {
        return isValid(root, Int.min, Int.max)
    }

    func isValid(_ node: TreeNode?, _ lower: Int, _ upper: Int) -> Bool{
        guard let node = node else {
            return true
        }
        if node.val <= lower || node.val >= upper{
            return false
        }
        return isValid(node.left, lower, node.val) && isValid(node.right, node.val, upper)
    }
}