function isAnagram(s: string, t: string): boolean {
    if (s.length != t.length) {
        return false
    }
    let hashmap: Map<string, number> = new Map()
    let value = ""
    let count = 0
    for (value of s) {
        hashmap.set(value, (hashmap.get(value) ?? 0) + 1)
    }
    for (value of t) {
        count = hashmap.get(value) ?? 0
        if (count > 0) {
            hashmap.set(value, count - 1)
        } else {
            return false
        }
    }
    return true
};