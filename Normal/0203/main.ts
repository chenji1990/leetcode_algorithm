function removeElements(head: ListNode | null, val: number): ListNode | null {
    let res = new ListNode()
    const header = res
    while(head){
        if (head.val != val){
            res.next = head
            res = res.next
            head = head.next
            continue
        }
        head = head.next
    }
    res.next = null
    return header.next
};