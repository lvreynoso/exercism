package hamming

import ( "errors" )

func Distance(a, b string) (int, error) {
    // check if strings are equivalent
    if len(a) != len(b) {
        return -1, errors.New("Sequences are not of equal length. The Distance function only works on sequences of the same length.")
    }

    // count the differences
    count := 0
    for i := 0; i < len(a); i++ {
        if a[i] != b[i] {
            count += 1
        }
    }

    return count, nil
}
