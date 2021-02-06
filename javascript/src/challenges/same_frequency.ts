// sameFrequency takes two numbers and comparer
// if the frequency of each digit is same with the other O(n)
function sameFrequency(num1: number, num2: number): boolean {
    // creates variables for storing a array of dig for both numbers
    const num1Dig = String(num1).split('')
    const num2Dig = String(num2).split('')
    // check if the length is different and return false
    if (num1Dig.length !== num2Dig.length) return false

    // creates a object to save the frequency of the first number
    const num1Freq: { [x: string]: number } = {}

    // loop thru the the first number and save the frequency in the object
    for (let i = 0; i < num1Dig.length; i++) {
        const dig = num1Dig[i]
        num1Freq[dig] = num1Freq[dig] + 1 || 1
    }

    // loop thru the the second number
    for (let i = 0; i < num2Dig.length; i++) {
        const dig = num2Dig[i]
        // if the dig is not on frequency object, return false
        if (!num1Freq[dig]) {
            return false
            // otherwise if the frequency is 0 or less return false
        } else if (num1Freq[dig] <= 0) {
            return false
            // otherwise decrease the frequency
        } else {
            num1Freq[dig] -= 1
        }
    }

    // return true if checks pass
    return true
}

export default sameFrequency
