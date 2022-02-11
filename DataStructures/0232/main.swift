class MyQueue {
    var stack: [Int] = []
    init() {

    }
    
    func push(_ x: Int) {
        self.stack.append(x)
    }
    
    func pop() -> Int {
        let length = self.stack.count
        if length <= 0 {
            return 0
        }
        return self.stack.removeFirst()
    }
    
    func peek() -> Int {
        return self.stack.first ?? 0
    }
    
    func empty() -> Bool {
        return self.stack.count == 0
    }
}
