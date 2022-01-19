class Solution {
    func setZeroes(_ matrix: inout [[Int]]) {
        if matrix.count == 0 || matrix[0].count == 0{
            return
        }
        let row = matrix.count
        let column = matrix[0].count

        var array: [[Int: Bool]] = [[:], [:]]
        for i in 0..<row{
            for j in 0..<column{
                if matrix[i][j] == 0{
                    array[0][i] = true
                    array[1][j] = true
                }
            }
        }

        for i in array[0].keys{
            for index in 0..<column{
                matrix[i][index] = 0
            }
        }
        for i in array[1].keys{
            for index in 0..<row{
                matrix[index][i] = 0
            }
        }
    }
}