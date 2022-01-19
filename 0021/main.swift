
  public class ListNode {
      public var val: Int
      public var next: ListNode?
      public init() { self.val = 0; self.next = nil; }
      public init(_ val: Int) { self.val = val; self.next = nil; }
      public init(_ val: Int, _ next: ListNode?) { self.val = val; self.next = next; }
  }
 
class Solution {
    func mergeTwoLists(_ list1: ListNode?, _ list2: ListNode?) -> ListNode? {
        var list1: ListNode? = list1
        var list2: ListNode? = list2

        var res: ListNode? = ListNode()
        let header = res
        while list1 != nil && list2 != nil{
            if (list1?.val ?? 0) < (list2?.val ?? 0){
                res?.next = list1
                list1 = list1?.next
            } else {
                res?.next = list2
                list2 = list2?.next
            }
            res = res?.next
        }

        if list1 == nil{
            res?.next = list2
        } else {
            res?.next = list1
        }
        return header?.next
    }
}