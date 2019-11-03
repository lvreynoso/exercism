// luhn.go
package luhn

import (
    "strconv"
    "fmt"
    "strings"
)

// Valid - given a string, will check if the string is valid under the Luhn algorithm.
func Valid(given string) (bool) {
    // strip spaces
    input := strings.Replace(given, " ", "", -1)

    // strings of length 1 are invalid
    if len(input) <= 1 {
        fmt.Println("Input length too small")
        return false
    }
    
    transformed := ""

    // first double every second digit (mod 9), starting from the right
    for i := 0; i < len(input); i++ {
        reverse_i := len(input) - 1 - i
        current, err := strconv.Atoi(string(input[reverse_i]))
        // check if the character is a digit, if not the sequence is not valid
        if err != nil {
            fmt.Println("Found a character that is not a digit")
            return false
        } else if i % 2 == 1 {
            double := current * 2

            // properly modulo the number so 9s remain a 9
            var newDigit int
            if double == 0 {
                newDigit = 0
            } else {
                moduloed := double % 9
                if moduloed == 0 {
                    newDigit = 9
                } else {
                    newDigit = moduloed
                }
            }

            transformed = strconv.Itoa(newDigit) + transformed
        } else {
            transformed = strconv.Itoa(current) + transformed
        }
    }

    // now sum up the transformed sequence
    sum := 0
    for _, char := range transformed {
        number, _ := strconv.Atoi(string(char))
        sum += number
    }
    if sum % 10 == 0 {
        return true
    } else {
        fmt.Printf("Sum did not add up to a multiple of 10. Instead added up %s to get %d.\n", transformed, sum)
        return false
    }
}