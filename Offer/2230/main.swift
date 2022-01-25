class MinStack {
    var stack = [Int]()
    var minStack = [Int]()
    
    init() {}
    
    func push(_ x: Int) {
        if stack.isEmpty{
            minStack.append(x)
        } else {
            minStack.append(Swift.min(x, (minStack.last ?? 0)))
        }
        stack.append(x)
    }
    
    func pop() {
        stack.removeLast()
        minStack.removeLast()
    }
    
    func top() -> Int {
        return stack.last ?? 0
    }
    
    func min() -> Int {
        return minStack.last ?? 0
    }
}