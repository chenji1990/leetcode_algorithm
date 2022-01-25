class CQueue {
    stack1: number[] = []
    stack2: number[] = []
    constructor() {}

    appendTail(value: number): void {
        this.stack1.push(value)
    }

    deleteHead(): number {
        if (this.stack2.length == 0){
            while(this.stack1.length > 0){
                this.stack2.push(this.stack1.pop() ?? 0)
            }
        }
        if (this.stack2.length > 0){
            return this.stack2.pop() ?? 0
        }
        return -1
    }
}