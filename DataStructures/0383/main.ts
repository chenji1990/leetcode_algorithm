function canConstruct(ransomNote: string, magazine: string): boolean {
    if (ransomNote.length > magazine.length){
        return false
    }
    let hashmap: Map<string, number> = new Map()
    let value = ""
    let count = 0
    for (value of magazine){
        hashmap.set(value, (hashmap.get(value) ?? 0) + 1)
    }

    for(value of ransomNote){
        count = hashmap.get(value) ?? 0
        if (count > 0){
            hashmap.set(value, count - 1)
        } else {
            return false
        }
    }

    return true
};