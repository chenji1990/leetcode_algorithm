class Solution {
    func isValid(_ s: String) -> Bool {
        let length = s.count
        if length == 0 || length % 2 != 0{
            return false
        }

        let hashmap: [Character: Character] = [
            ")": "(",
            "]": "[",
            "}": "{",
        ]

        var stack = [Character]()
        var lastIndex = 0

        for char in s{
            lastIndex = stack.count - 1
            if lastIndex <= -1 || stack[lastIndex] != (hashmap[char] ?? "_"){
                stack.append(char)
                continue
            }
            stack.removeLast()
        }
        return stack.count == 0
    }
}