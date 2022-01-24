function getDecimalValue(head: ListNode | null): number {
    let sum = 0
    while (head){
        sum = sum * 2 + head.val
        head = head.next
    }
    return sum
};