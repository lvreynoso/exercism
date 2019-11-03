// package twofer has just one function, ShareWith.
package twofer

// package twofer takes in a string and uses that string as a name, X, in the sentence
// "One for X, one for me." If no string is passed, it uses "you" in its place.
func ShareWith(name string) string {
    outputName := ""
    if (len(name) > 0) {
        outputName = name
    } else {
        outputName = "you"
    }
	return "One for " + outputName + ", one for me."
}
