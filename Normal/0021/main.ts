function mergeTwoLists(list1: ListNode | null, list2: ListNode | null): ListNode | null {
    let res = new ListNode()
    const header = res
    while(list1 && list2){
        if (list1.val < list2.val){
            res.next = list1
            list1 = list1.next
        } else {
            res.next = list2
            list2 = list2.next
        }
        res = res.next
    }

    if (list1){
        res.next = list1
    } else {
        res.next = list2
    }
    return header.next
};