function twoSum(nums: number[], target: number): number[] {
    let hashmap: Map<number, number> = new Map()
    let [value, i2] = [0, 0]
    for (let i1 in nums) {
        value = target - nums[i1]
        i2 = hashmap.get(value) ?? 0

        if (i2 != undefined) {
            return [i2, Number(i1)]
        }
        hashmap.set(nums[i1], Number(i1))
    }
    return []
};
