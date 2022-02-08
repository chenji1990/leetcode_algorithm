class Solution {
    func isValidSudoku(_ board: [[Character]]) -> Bool {
        var array = Array.init(repeating: [
            [0, 0, 0, 0, 0, 0, 0, 0, 0],
            [0, 0, 0, 0, 0, 0, 0, 0, 0],
            [0, 0, 0, 0, 0, 0, 0, 0, 0],
        ], count: 9)

        var index: Int = 0
        var section: Int = 0
        for i in 0..<9{
            for j in 0..<9{
                if board[i][j] == "."{
                    continue
                }
                index = Int((board[i][j].asciiValue ?? 0) - 49)
                if array[index][0][i] > 0 || array[index][1][j] > 0{
                    return false
                }
                section = (i / 3) * 3 + (j / 3)
                if array[index][2][section] > 0{
                    return false
                }
                array[index][0][i] = 1
                array[index][1][j] = 1
                array[index][2][section] = 1
            }
        }
        return true
    }
}