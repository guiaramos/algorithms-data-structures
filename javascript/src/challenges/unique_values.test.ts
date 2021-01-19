import { countUniqueValues } from './unique_values'

describe('Count Unique Values Challenge', () => {
    it('should return 0 when array is empty', () => {
        const emptyArray: [] = []
        const got = countUniqueValues(emptyArray)
        const want = 0
        expect(got).toBe(want)
    })

    it('should return 1 when array has 1 unique value', () => {
        const values: number[] = [1]
        const got = countUniqueValues(values)
        const want = 1
        expect(got).toBe(want)
    })

    it('should return 2 when array has 2 unique values', () => {
        const values: number[] = [1, 2]
        const got = countUniqueValues(values)
        const want = 2
        expect(got).toBe(want)
    })

    it('should return 2 when array has 2 unique values with repeated numbers', () => {
        const values: number[] = [1, 1, 1, 1, 1, 2]
        const got = countUniqueValues(values)
        const want = 2
        expect(got).toBe(want)
    })

    it('should return 4 when array has 4 unique values', () => {
        const values: number[] = [0, 1, 2, 3]
        const got = countUniqueValues(values)
        const want = 4
        expect(got).toBe(want)
    })

    it('should return 7 when array has 7 unique values and repeated numbers', () => {
        const values: number[] = [1, 2, 3, 4, 4, 4, 7, 7, 12, 12, 13]
        const got = countUniqueValues(values)
        const want = 7
        expect(got).toBe(want)
    })

    it('should return 4 when array has 4 unique values with negative values', () => {
        const values: number[] = [-2, -1, -1, 0, 1]
        const got = countUniqueValues(values)
        const want = 4
        expect(got).toBe(want)
    })
})
