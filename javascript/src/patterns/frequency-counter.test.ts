import { sameArraySquared } from './frequency-counter'

describe('Frequency Counter Pattern', () => {
    it('should return false if length of the arrays are different', () => {
        const firstArray: number[] = [1, 2, 4, 5]
        const secondArray: number[] = [1, 2, 4, 5, 6]

        const isSame = sameArraySquared(firstArray, secondArray)

        expect(isSame).toBeFalsy()
    })

    it('should return false if number is not square', () => {
        const firstArray: number[] = [1, 2, 3]
        const secondArray: number[] = [1, 4, 6]

        const isSame = sameArraySquared(firstArray, secondArray)

        expect(isSame).toBeFalsy()
    })

    it('should return true for ordered array', () => {
        const firstArray: number[] = [1, 2, 3]
        const secondArray: number[] = [1, 4, 9]

        const isSame = sameArraySquared(firstArray, secondArray)

        expect(isSame).toBeTruthy()
    })

    it('should return true for unordered array', () => {
        const firstArray: number[] = [2, 3, 1]
        const secondArray: number[] = [9, 1, 4]

        const isSame = sameArraySquared(firstArray, secondArray)

        expect(isSame).toBeTruthy()
    })
})
