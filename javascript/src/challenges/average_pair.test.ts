import averagePair from './average_pair'

describe('Average Pair Challenge', () => {
    it('should return false when array is empty', () => {
        const intArr: number[] = []
        const average = 4
        const isAverage = averagePair(intArr, average)
        expect(isAverage).toBeFalsy()
    })

    it('should return true with small numbers array', () => {
        const intArr: number[] = [1, 2, 3]
        const average = 2.5
        const isAverage = averagePair(intArr, average)
        expect(isAverage).toBeTruthy()
    })
    it('should return true with long numbers array', () => {
        const intArr: number[] = [1, 3, 3, 5, 6, 7, 10, 12, 19]
        const average = 8
        const isAverage = averagePair(intArr, average)
        expect(isAverage).toBeTruthy()
    })
    it('should return false with long numbers array and negative', () => {
        const intArr: number[] = [-1, 0, 3, 4, 5, 6]
        const average = 4.1
        const isAverage = averagePair(intArr, average)
        expect(isAverage).toBeFalsy()
    })
})
