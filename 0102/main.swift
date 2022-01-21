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
    func levelOrder(_ root: TreeNode?) -> [[Int]] {
        guard let root = root else {
            return []
        }
        var array = [[Int]]()
        var row: [TreeNode] = [root]
        var nodeArray = [TreeNode]()
        var valueArray = [Int]()
        while row.count != 0{
            nodeArray = []
            valueArray = []
            for node in row{
                valueArray.append(node.val)
                if let left = node.left{
                    nodeArray.append(left)
                }
                if let right = node.right{
                    nodeArray.append(right)
                }
            }
            array.append(valueArray)
            row = nodeArray
        }
        return array
    }
}