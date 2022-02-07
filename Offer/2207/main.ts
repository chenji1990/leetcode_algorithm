function reversePrint(head: ListNode | null): number[] {
    let stack: number[] = []
    while(head != null){
        stack.push(head.val)
        head = head.next
    }
    let res: number[] = []
    while(stack.length > 0){
        res.push(stack.pop() ?? 0)
    }
    return res
};