function setZeroes(matrix: number[][]): void {
    if (matrix.length == 0 || matrix[0].length == 0) {
        return
    }
    let [row, column] = [matrix.length, matrix[0].length]
    let [i, j, index] = [0, 0, 0]
    let array: Map<number, boolean>[] = [new Map(), new Map()]

    for (i = 0; i < row; i++) {
        for (j = 0; j < column; j++) {
            if (matrix[i][j] == 0) {
                array[0][i] = true
                array[1][j] = true
            }
        }
    }

    for (let i in array[0]) {
        for (index = 0; index < column; index++) {
            matrix[Number(i)][index] = 0
        }
    }

    for (let i in array[1]) {
        for (index = 0; index < row; index++) {
            matrix[index][Number(i)] = 0
        }
    }
};