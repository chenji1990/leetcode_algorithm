class Solution {
    func firstUniqChar(_ s: String) -> Int {
        var hashmap = [Character: Int]()

        for (index, value) in s.enumerated(){
            if let _ = hashmap[value]{
                hashmap[value] = -1
            } else {
                hashmap[value] = index
            }
        }

        var minIndex: Int = Int.max
        for (_, index) in hashmap{
            if index != -1 && index < minIndex{
                minIndex = index
            }
        }
        if minIndex == Int.max{
            return -1
        }
        return minIndex
    }
}