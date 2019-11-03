// gigasecond.h

#ifndef GIGASECOND_H
#define GIGASECOND_H
#define EXERCISM_RUN_ALL_TESTS

#include "boost/date_time/posix_time/posix_time.hpp"

using namespace boost::posix_time;

namespace gigasecond {

    ptime advance(ptime input);

}

#endif