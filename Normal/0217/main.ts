function containsDuplicate(nums: number[]): boolean {
    let hashmap = {}
    for (let value of nums){
        if (hashmap[value] === true){
            return true
        }
        hashmap[value] = true
    }
    return false
}