public class ListNode {
    public var val: Int
    public var next: ListNode?
    public init(_ val: Int) {
        self.val = val
        self.next = nil
    }
}

class Solution {
    func reverseList(_ head: ListNode?) -> ListNode? {
        var head = head
        var res = ListNode(0)
        var temp: ListNode? = nil
        while head != nil{
            temp = head?.next
            head?.next = res.next
            res.next = head
            head = temp
        }
        return res.next
    }
}