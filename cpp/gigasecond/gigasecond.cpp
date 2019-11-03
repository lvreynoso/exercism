// gigasecond.cpp

#include "boost/date_time/posix_time/posix_time.hpp"

using namespace boost::posix_time;

namespace gigasecond {

    ptime advance(ptime input) {
        ptime output = input + seconds(1000000000);
        return output;
    }

}