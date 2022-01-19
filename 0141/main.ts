
  class ListNode {
      val: number
      next: ListNode | null
      constructor(val?: number, next?: ListNode | null) {
          this.val = (val===undefined ? 0 : val)
          this.next = (next===undefined ? null : next)
      }
}

 function hasCycle(head: ListNode | null): boolean {
    let hashmap: Map<ListNode, boolean> = new Map()
    while (head != null){
        if (hashmap.get(head)){
            return true
        }
        hashmap.set(head, true)
        head = head.next
    }
    return false
};