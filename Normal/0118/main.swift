class Solution {
    func generate(_ numRows: Int) -> [[Int]] {
        if numRows <= 0{
            return []
        }
        if numRows == 1{
            return [[1]]
        }
        if numRows == 2{
            return [[1], [1, 1]]
        }
        var res: [[Int]] = [[1], [1, 1]]
        var rowArray: [Int] = [1]
        
        for row in 2..<numRows{
            for column in 1..<row{
                rowArray.append(res[row - 1][column - 1] + res[row - 1][column])
            }
            rowArray.append(1)
            res.append(rowArray)
            rowArray = [1]
        }
        return res
    }
}