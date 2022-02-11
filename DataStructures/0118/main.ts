function generate(numRows: number): number[][] {
    if (numRows <= 0){
        return []
    }
    if (numRows == 1){
        return [[1]]
    }
    if (numRows == 2){
        return [[1], [1, 1]]
    }
    let res: number[][] = [[1], [1, 1]]
    let rowArray: number[] = [1]
    let row = 0
    let column = 0
    for (row = 2; row < numRows; row++){
        for (column = 1; column < row; column++){
            rowArray.push(res[row - 1][column - 1] + res[row - 1][column])
        }
        rowArray.push(1)
        res.push(rowArray)
        rowArray = [1]
    }
    return res
};