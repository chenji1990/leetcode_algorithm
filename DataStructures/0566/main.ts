function matrixReshape(mat: number[][], r: number, c: number): number[][] {
    if (mat.length == 0 || mat[0].length == 0 || r * c != mat.length * mat[0].length) {
        return mat
    }

    var res: number[][] = []
    var newRowArray: number[] = []
    var column = 0
    var rowArray: number[] = []
    var value = 0
    for (rowArray of mat) {
        for (value of rowArray) {
            if (column < c) {
                newRowArray.push(value)
                column += 1
            } else {
                res.push(newRowArray)
                newRowArray = [value]
                column = 1
            }
        }
    }
    res.push(newRowArray)
    return res
};