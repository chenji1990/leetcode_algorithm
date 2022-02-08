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
    func lowestCommonAncestor(_ root: TreeNode?, _ p: TreeNode?, _ q: TreeNode?) -> TreeNode? {
        var root = root
        guard var p = p, var q = q else{
            return root
        }
        if p.val > q.val{
            let temp = p
            p = q
            q = temp
        }
        while root != nil{
            if root!.val > q.val{
                root = root!.left
                continue
            }
            if root!.val < p.val{
                root = root!.right
                continue
            }
            return root
        }
        return root
    }
}