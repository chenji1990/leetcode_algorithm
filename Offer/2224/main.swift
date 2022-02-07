  public class ListNode {
      public var val: Int
      public var next: ListNode?
      public init() { self.val = 0; self.next = nil; }
      public init(_ val: Int) { self.val = val; self.next = nil; }
      public init(_ val: Int, _ next: ListNode?) { self.val = val; self.next = next; }
  }

class Solution {
    func reverseList(_ head: ListNode?) -> ListNode? {
        var head = head
        var root = ListNode()
        var temp: ListNode? = nil
        while head != nil{
            temp = head!.next
            head!.next = root.next
            root.next = head
            head = temp
        }
        return root.next
    }
}