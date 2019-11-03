// leap.cpp

namespace leap {

    bool is_leap_year(int year) {
        bool result = false;

        if (year % 400 == 0) {
            result = true;
        } else if (year % 4 == 0 && year % 100 != 0) {
            result = true;
        }

        return result;
    }

}