class MyQueue {

    stack: number[] = []

    constructor() {

    }

    push(x: number): void {
        this.stack.push(x)
    }

    pop(): number {
        let length = this.stack.length
        if (length <= 0){
            return 0
        }
        let firstValue = this.stack[0]
        this.stack = this.stack.slice(1, length)
        return firstValue
    }

    peek(): number {
        if (this.stack.length == 0){
            return 0
        }
        return this.stack[0]
    }

    empty(): boolean {
        return this.stack.length == 0
    }
}