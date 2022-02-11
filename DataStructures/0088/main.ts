function merge(nums1: number[], m: number, nums2: number[], n: number): void {
    let [p1, p2] = [m - 1, n - 1]
    let [value1, value2] = [0, 0]

    for (let i = m + n - 1; i >=0; i--) {
        if (p1 < 0){
            nums1[i] = nums2[p2]
            p2--
            continue
        }
        if (p2 < 0){
            nums1[i] = nums1[p1]
            p1--
            continue
        }
        [value1, value2] = [nums1[p1], nums2[p2]]
        if (value2 >= value1){
            nums1[i] = nums2[p2]
            p2--
        } else {
            nums1[i] = nums1[p1]
            p1--
        }
    }
};