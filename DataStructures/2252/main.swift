class Node{
    var val: Int = 0
    var next: Node?
    var random: Node?

    init(_ val: Int, _ next: Node? = nil, _ random: Node? = nil){
        self.val = val
        self.next = next
        self.random = random
    }
}

class Solution{
    func copyRandomList(head: Node?) -> Node?{
        guard let head = head else {
            return nil
        }

        let root = Node(0)
        var tempRoot: Node = root
        var tempHead: Node? = head
        var i = 0

        var map1 = [Node: Int]()
        var map2 = [Int: Node]()

        while tempHead != nil{
            tempRoot.next = Node(tempHead!.val)
            map1[tempHead!] = i
            map2[i] = tempRoot.next!

            tempRoot = tempRoot.next
            tempHead = tempHead!.next

            i += 1
        }

        var tempRandom: Node? = nil
        tempRoot = root
        tempHead = head

        while tempHead != nil{
            if let tempRandom = tempHead?.random,
                let index = map1[tempRandom]{
                    tempRoot.next?.random = map2[index]
            }
            tempRoot = tempRoot.next
            tempHead = tempHead?.next
        }

        return root.next
    }
}
    