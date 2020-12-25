// validAnagram checks if two words is an anagram O(n)
function validAnagram (firstString: string, secondString: string): Boolean {
  // Check if the length of both strings are same
  if (firstString.length !== secondString.length) { return false }

  // Create objects for the both strings to save the frequency
  const frequencyObj: Record<string, number> = {}

  // Iterate thru first string and save the frequency on object
  for (const letter of firstString) {
    frequencyObj[letter] = (frequencyObj[letter] | 0) + 1
  }

  // Iterate thru te first frequency object
  for (const letter of secondString) {
    // Check if the the letter exists on second frequency object
    if (!frequencyObj[letter]) { return false } else { frequencyObj[letter] -= 1 }
  }

  return true
}

export { validAnagram }
