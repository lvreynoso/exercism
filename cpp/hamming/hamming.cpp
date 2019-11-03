// hamming.cpp

#include <string>

using namespace std;

namespace hamming {

    int compute(string first, string second) {

        // first check if string lengths are equal
        if (first.length() != second.length()) {
            throw domain_error("DNA Sequences of unequal length cannot be used to calculate a Hamming distance.");
        }

        int distance = 0;

        // compare strings
        for (string::size_type index = 0; index < first.length(); index++) {
            if (first[index] != second[index]) {
                distance += 1;
            }
        }

        return distance;
    }

}