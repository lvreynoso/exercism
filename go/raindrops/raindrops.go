// raindrops.go

package raindrops

import ( "strconv" )

func Convert(inputNumber int) (string) {
    outputString := ""

    // check factors
    if inputNumber % 3 == 0 {
        outputString += "Pling"
    }
    if inputNumber % 5 == 0 {
        outputString += "Plang"
    }
    if inputNumber % 7 == 0 {
        outputString += "Plong"
    }

    // if none of them apply, output the number
    if len(outputString) == 0 {
        outputString += strconv.Itoa(inputNumber)
    }

    return outputString
}