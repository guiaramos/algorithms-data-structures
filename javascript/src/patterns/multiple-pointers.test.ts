import { sumZero } from './multiple-pointers'

describe('Multiple Pointers Pattern', () => {
    it('should return undefined on empty array', () => {
        const got = sumZero([])
        expect(got).toBeUndefined()
    })

    it('should find the first pair where sum is zero', () => {
        const test = [-16, -2, -1, 0, 1, 2, 16]
        const got = sumZero(test)
        const want = [-16, 16]
        expect(got).toMatchObject(want)
    })

    it('should find in any location pair where sum is zero', () => {
        const test = [-16, -8, -2, 0, 2, 5, 17]
        const got = sumZero(test)
        const want = [-2, 2]
        expect(got).toMatchObject(want)
    })
})
