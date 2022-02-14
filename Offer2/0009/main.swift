class CQueue {
    var stack1 = [Int]()
    var stack2 = [Int]()
    init() {}
    
    func appendTail(_ value: Int) {
        self.stack1.append(value)
    }
    
    func deleteHead() -> Int {
        if self.stack2.isEmpty{
            while !self.stack1.isEmpty{
                self.stack2.append(self.stack1.removeLast())
            }
        }
        if !self.stack2.isEmpty{
            return self.stack2.removeLast()
        }
        return -1
    }
}
