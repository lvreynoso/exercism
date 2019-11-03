// reverse_string.cpp

#include <string>

using namespace std;

namespace reverse_string {

    string reverse_string(string input) {

        string output = "";

        for (string::size_type i = 0; i < input.length(); i++) {
            string::size_type head = input.length() - 1 - i;
            output += input[head];
        }

        return output;

    }

}