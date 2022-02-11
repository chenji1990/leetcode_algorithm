class Solution {
    func getDecimalValue(_ head: ListNode?) -> Int {
        var head = head
        var sum = 0
        while head != nil{
            sum = sum * 2 + head!.val
            head = head!.next
        }
        return sum
    }
}