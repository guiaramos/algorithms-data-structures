import sameFrequency from './same_frequency'

describe('Same Frequency Challenge', () => {
    it('should return false when input has diff length', () => {
        const num1 = 22
        const num2 = 222
        const result = sameFrequency(num1, num2)
        expect(result).toBeFalsy()
    })

    it('should return false when length is same but frequency is wrong', () => {
        const num1 = 34
        const num2 = 14
        const result = sameFrequency(num1, num2)
        expect(result).toBeFalsy()
    })

    it('should return true when length and frequency is same with small numbers', () => {
        const num1 = 182
        const num2 = 281
        const result = sameFrequency(num1, num2)
        expect(result).toBeTruthy()
    })

    it('should return true when length and frequency is same with large numbers', () => {
        const num1 = 3589578
        const num2 = 5879385
        const result = sameFrequency(num1, num2)
        expect(result).toBeTruthy()
    })
})
