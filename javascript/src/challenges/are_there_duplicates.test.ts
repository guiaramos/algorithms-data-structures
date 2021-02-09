import areThereDuplicates from './are_there_duplicates'

describe('Are There Duplicates Challenge', () => {
    it('should return false when the int args are not repeated', () => {
        const isDuplicate = areThereDuplicates(1, 2, 3)
        expect(isDuplicate).toBeFalsy()
    })
    it('should return true when the int args are repeated', () => {
        const isDuplicate = areThereDuplicates(1, 2, 2)
        expect(isDuplicate).toBeTruthy()
    })
    it('should return true when the int args are unsorted repeated', () => {
        const isDuplicate = areThereDuplicates(2, 1, 2)
        expect(isDuplicate).toBeTruthy()
    })
    it('should return true when the string args are repeated', () => {
        const isDuplicate = areThereDuplicates('a', 'b', 'c', 'a')
        expect(isDuplicate).toBeTruthy()
    })
    it('should return false when the string args are not repeated', () => {
        const isDuplicate = areThereDuplicates('a', 'b', 'c')
        expect(isDuplicate).toBeFalsy()
    })
})
