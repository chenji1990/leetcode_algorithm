
func containsDuplicate(_ nums: [Int]) -> Bool {
    var hashmap = [Int: Bool]()
    for value in nums{
        if let res = hashmap[value], res{
            return true
        }
        hashmap[value] = true
    }
    return false
}

