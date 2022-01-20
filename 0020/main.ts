function isValid(s: string): boolean {
    let length = s.length
    if (length == 0 || length % 2 != 0){
        return false
    }
    let hashmap: Map<string, string> = new Map([
        [")", "("],
        ["]", "["],
        ["}", "{"],
    ])
    let stack: string[] = []
    let lastIndex = 0
    let char = ""
    for (char of s){
        lastIndex = stack.length - 1
        if (lastIndex <= -1 || stack[lastIndex] !== hashmap.get(char)){
            stack.push(char)
            continue
        }
        stack = stack.slice(0, lastIndex)
    }
    return stack.length == 0
};

console.log(isValid("()"))