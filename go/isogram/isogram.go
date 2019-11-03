// isogram.go

package isogram

import (
    "unicode"
)

// IsIsogram takes in a string and determines whether or not
// that string is an isogram, returning a boolean.
func IsIsogram(inputWord string) bool {

    result := true
    letterTracker := make(map[rune]bool)

    for _, character := range inputWord {

        current := unicode.ToUpper(character)

        if letterTracker[current] {
            result = false
        }

        if unicode.IsLetter(current) {
            letterTracker[current] = true
        }
    }

    return result
}
