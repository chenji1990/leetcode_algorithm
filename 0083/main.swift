public class ListNode {
    public var val: Int
    public var next: ListNode?
    public init(_ val: Int) {
        self.val = val
        self.next = nil
    }
}

class Solution {
    func deleteDuplicates(_ head: ListNode?) -> ListNode? {
        var head = head
        var res = ListNode(Int.min)
        var header = res
        while head != nil{
            if head!.val != res.val{
                res.next = head
                head = head?.next
                res = res.next!
                continue
            }
            head = head?.next
        }
        res.next = nil
        return header.next
    }
}