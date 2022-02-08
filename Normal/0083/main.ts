function deleteDuplicates(head: ListNode | null): ListNode | null {
    let res = new ListNode(-1e9)
    let header = res
    while (head){
        if (head.val != res.val){
            res.next = head
            head = head.next
            res = res.next
            continue
        }
        head = head.next
    }
    res.next = null
    return header.next
};