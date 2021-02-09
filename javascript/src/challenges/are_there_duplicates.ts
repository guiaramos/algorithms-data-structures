// areThereDuplicates accepts many args and compare if there are duplicates args
// O(n log n) when using Set
// O(n) or using loop
function areThereDuplicates(...args: number[] | string[]): boolean {
    // Solution O(n log n)
    // return new Set([...args]).size !== args.length

    // sort the args
    args.sort((a: any, b: any) => (a > b ? 1 : -1))

    // create two pointers for the start and next pos of the arr
    let start = 0
    let next = 1

    // loop thru the array
    while (next < args.length) {
        // if the values are same, returns true
        if (args[start] === args[next]) {
            return true
        }
        // increases the pointers
        start++
        next++
    }

    // if the check passes return false
    return false
}

export default areThereDuplicates
