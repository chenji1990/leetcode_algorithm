public class ListNode {
    public var val: Int
    public var next: ListNode?
    public init() { self.val = 0; self.next = nil; }
    public init(_ val: Int) { self.val = val; self.next = nil; }
    public init(_ val: Int, _ next: ListNode?) { self.val = val; self.next = next; }
}

class Solution {
    func reversePrint(_ head: ListNode?) -> [Int] {
        var head = head
        var stack = [Int]()
        while head != nil{
            stack.append(head!.val)
            head = head!.next
        }
        var res = [Int]()
        while !stack.isEmpty{
            res.append(stack.removeLast())
        }
        return res
    }
}