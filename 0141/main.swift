public class ListNode {
    public var val: Int
    public var next: ListNode?
    public init(_ val: Int) {
        self.val = val
        self.next = nil
    }
}

class Solution {
    func hasCycle(_ head: ListNode?) -> Bool {
        var head = head
        var hashmap = [String: Bool]()
        var addr = ""
        while head != nil{
            addr = Unmanaged.passUnretained(head!).toOpaque().debugDescription
            if let res = hashmap[addr], res{
                return true
            }
            hashmap[addr] = true
            head = head?.next
        }
        return false
    }
}