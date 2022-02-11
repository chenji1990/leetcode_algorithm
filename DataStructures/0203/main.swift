import Foundation

  public class ListNode {
      public var val: Int
      public var next: ListNode?
      public init() { self.val = 0; self.next = nil; }
      public init(_ val: Int) { self.val = val; self.next = nil; }
      public init(_ val: Int, _ next: ListNode?) { self.val = val; self.next = next; }
  }
 
class Solution {
    func removeElements(_ head: ListNode?, _ val: Int) -> ListNode? {
        var head = head
        var res: ListNode? = ListNode()
        let header = res
        while head != nil{
            if head!.val != val{
                res?.next = head
                res = res?.next
                head = head!.next
                continue
            }
            head = head!.next
        }
        res?.next = nil
        return header?.next
    }
}