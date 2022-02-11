class Solution {
    func isAnagram(_ s: String, _ t: String) -> Bool {
        if s.count != t.count{
            return false
        }
        var hashmap = [Character: Int]()
        var count = 0
        for value in s{
            hashmap[value] = (hashmap[value] ?? 0) + 1
        }
        for value in t{
            count = hashmap[value] ?? 0
            if count > 0{
                hashmap[value] = count - 1
            } else {
                return false
            }
        }

        return true
    }
}