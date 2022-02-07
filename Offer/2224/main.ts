function reverseList(head: ListNode | null): ListNode | null {
    let root = new ListNode()
    let temp: ListNode | null = null
    while (head != null){
        temp = head.next
        head.next = root.next
        root.next = head
        head = temp
    }
    return root.next
};