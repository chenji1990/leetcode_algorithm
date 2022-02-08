function intersect(nums1: number[], nums2: number[]): number[] {
    if (nums1.length > nums2.length){
        return intersect(nums2, nums1)
    }

    let hashmap = new Map()
    let value = 0
    let count = 0
    for (value of nums1){
        count = hashmap[value]
        if (count != undefined && count > 0){
            hashmap[value] = count + 1
        } else {
            hashmap[value] = 1
        }
    }

    let res: number[] = []
    for (value of nums2){
        count = hashmap[value]
        if (count != undefined && count > 0){
            res.push(value)
            hashmap[value] = count - 1
        }
    }
    return res
};