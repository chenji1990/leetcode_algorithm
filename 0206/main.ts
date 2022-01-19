function reverseList(head: ListNode | null): ListNode | null {
    let res = new ListNode()
    let temp: ListNode | null = null
    while (head){
        temp = head.next
        head.next = res.next
        res.next = head
        head = temp
    }
    return res.next
};