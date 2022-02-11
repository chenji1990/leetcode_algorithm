function isValidSudoku(board: string[][]): boolean {
    let array: number[][][] = []
    let [index, i, j, section] = [0, 0, 0, 0]
    for (index = 0; index < 9; index++) {
        array.push([
            [0, 0, 0, 0, 0, 0, 0, 0, 0],
            [0, 0, 0, 0, 0, 0, 0, 0, 0],
            [0, 0, 0, 0, 0, 0, 0, 0, 0],
        ])
    }

    for(i = 0; i < 9; i++){
        for(j = 0; j < 9; j++){
            if (board[i][j] == "."){
                continue
            }
            index = board[i][j].charCodeAt(0) - 49
            if (array[index][0][i] > 0 || array[index][1][j] > 0){
                return false
            } 
            section = Math.floor(i / 3) * 3 + Math.floor(j / 3)
            if (array[index][2][section] > 0){
                return false
            }

            array[index][0][i] = 1
            array[index][1][j] = 1
            array[index][2][section] = 1
        }
    }
    return true
};