function firstUniqChar(s: string): number {
    let hashmap: Map<string, number> = new Map()
    let index = 0
    let value = ""
    let len = s.length

    for(index = 0; index < len; index++){
        value = s[index]
        if (hashmap.get(value) != undefined){
            hashmap.set(value, -1)
        } else {
            hashmap.set(value, index)
        }
    }

    let minIndex = 1e9

    hashmap.forEach((index) => {
        if (index != -1 && index < minIndex){
            minIndex = index
        }
    })

    if (minIndex == 1e9){
        return -1
    }

    return minIndex
};