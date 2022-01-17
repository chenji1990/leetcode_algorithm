class Solution {
    func matrixReshape(_ mat: [[Int]], _ r: Int, _ c: Int) -> [[Int]] {
        if mat.count == 0 || mat[0].count == 0 || r * c != mat.count * mat[0].count{
            return mat
        }

        var res: [[Int]] = []
        var newRowArray: [Int] = []
        var column = 0

        for rowArray in mat{
            for value in rowArray{
                if column < c{
                    newRowArray.append(value)
                    column += 1
                } else {
                    res.append(newRowArray)
                    newRowArray = [value]
                    column = 1
                }
            }
        }
        res.append(newRowArray)
        return res
    }
}