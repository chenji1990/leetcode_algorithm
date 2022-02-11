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
    func isSymmetric(_ root: TreeNode?) -> Bool {
        return check(root, root)
    }

    func check(_ p: TreeNode?, _ q: TreeNode?) -> Bool{
        if p == nil && q == nil{
            return true
        }
        if p == nil || q == nil{
            return false
        }
        return p!.val == q!.val && check(p?.left, q?.right) && check(p?.right, q?.left)
    }
}