class MinStack {
    stack: number[] = []
    minStack: number[] = []
    constructor() { }

    push(x: number): void {
        if (this.stack.length == 0) {
            this.minStack.push(x)
        } else {
            this.minStack.push(Math.min(x, this.minStack[this.minStack.length - 1]))
        }
        this.stack.push(x)
    }

    pop(): void {
        this.stack.pop()
        this.minStack.pop()
    }

    top(): number {
        if (this.stack.length == 0) {
            return 0
        }
        return this.stack[this.minStack.length - 1]
    }

    min(): number {
        if (this.minStack.length == 0) {
            return 0
        }
        return this.minStack[this.minStack.length - 1]
    }
}