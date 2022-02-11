class Solution {
    func canConstruct(_ ransomNote: String, _ magazine: String) -> Bool {
        if ransomNote.count > magazine.count{
            return false
        }
        var hashmap = [Character: Int]()
        for value in magazine{
            hashmap[value] = (hashmap[value] ?? 0) + 1
        }
        for value in ransomNote{
            if let count = hashmap[value], count > 0{
                hashmap[value] = count - 1
            } else {
                return false
            }
        }
        return true
    }
}