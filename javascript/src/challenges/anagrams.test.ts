import { validAnagram } from './anagrams'

describe('Anagrams Check Challenge', () => {
    it('should return true if empty string', () => {
        const firstString = ''
        const secondString = ''

        const isValid = validAnagram(firstString, secondString)

        expect(isValid).toBeTruthy()
    })

    it('should return false if the length is different', () => {
        const firstString = 'awesome'
        const secondString = 'awesom'

        const isValid = validAnagram(firstString, secondString)

        expect(isValid).toBeFalsy()
    })

    it('should return true if pass correct anagram', () => {
        const firstString = 'anagram'
        const secondString = 'nagaram'

        const isValid = validAnagram(firstString, secondString)

        expect(isValid).toBeTruthy()
    })

    it('should return false if the frequency of letters are not same', () => {
        const firstString = 'aaz'
        const secondString = 'zza'

        const isValid = validAnagram(firstString, secondString)

        expect(isValid).toBeFalsy()
    })

    it('should return false for "rat" and "car"', () => {
        const firstString = 'rat'
        const secondString = 'car'

        const isValid = validAnagram(firstString, secondString)

        expect(isValid).toBeFalsy()
    })

    it('should return true for "qwerty" and "qeywrt"', () => {
        const firstString = 'qwerty'
        const secondString = 'qeywrt'

        const isValid = validAnagram(firstString, secondString)

        expect(isValid).toBeTruthy()
    })

    it('should return true for "texttwisttime" and "timetwisttext"', () => {
        const firstString = 'texttwisttime'
        const secondString = 'timetwisttext'

        const isValid = validAnagram(firstString, secondString)

        expect(isValid).toBeTruthy()
    })
})
