import maxSubarraySum from './sliding_window'

describe('Sliding Window Pattern', () => {
    it('should return null when array is empty', function () {
        const test: number[] = []
        const num = 4
        const got = maxSubarraySum(test, num)

        expect(got).toBeNull()
    })

    it('should return 6 when array has max num of 4 same equals', function () {
        const test: number[] = [4, 2, 1, 6]
        const num = 1
        const got = maxSubarraySum(test, num)
        const want = 6

        expect(got).toBe(want)
    })

    it('should return 10 when array has max num of 2', function () {
        const test: number[] = [1, 2, 5, 2, 8, 1, 5]
        const num = 2
        const got = maxSubarraySum(test, num)
        const want = 10

        expect(got).toBe(want)
    })

    it('should return 17 when array has max num of 4', function () {
        const test: number[] = [1, 2, 5, 2, 8, 1, 5]
        const num = 4
        const got = maxSubarraySum(test, num)
        const want = 17

        expect(got).toBe(want)
    })

    it('should return 13 when array has max num of 4', function () {
        const test: number[] = [4, 2, 1, 6, 2]
        const num = 4
        const got = maxSubarraySum(test, num)
        const want = 13

        expect(got).toBe(want)
    })
})
