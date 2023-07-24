//
// rtcg.cpp - functionality common to rtcg tests
//
// Copyright (C) 2023 University of York and others
//
// SPDX-License-Identifier: MIT
//

#include <ostream>

#include "rtcg.h"

//
// Outcome
//

const char* rtcg::outstr(rtcg::Outcome o)
{
  switch (o)
  {
  case rtcg::Outcome::UNSET:
    return "unset";
  case rtcg::Outcome::INC:
    return "inconclusive";
  case rtcg::Outcome::PASS:
    return "passed";
  case rtcg::Outcome::FAIL:
    return "failed";
  }
  return "???";
}

std::ostream& operator<<(std::ostream& stm, rtcg::Outcome o)
{
  return stm << rtcg::outstr(o);
}

//
// Status
//

const char* rtcg::status::explain(rtcg::Status s)
{
  switch (s) {
  case rtcg::Status::WAIT_START:
    return "waiting to start";
  case rtcg::Status::WAIT_IN:
    return "waiting for input acknowledgement";
  case rtcg::Status::WAIT_OUT:
    return "waiting for output";
  case rtcg::Status::OFF_SCRIPT:
    return "an event happened that took the test off-script";
  case rtcg::Status::TIMEOUT:
    return "the test ran out of time";
  case rtcg::Status::FAIL:
    return "an unwanted behaviour occurred";
  case rtcg::Status::BUG:
    return "an internal error occurred";
  default:
    return "unknown (this should not occur)";
  }
}

std::ostream& operator<<(std::ostream& os, rtcg::PRINT_COLOR c)
{
  switch(c)
  {
    case rtcg::PRINT_COLOR::BLACK    : os << "\033[1;30m"; break;
    case rtcg::PRINT_COLOR::RED      : os << "\033[1;31m"; break;
    case rtcg::PRINT_COLOR::GREEN    : os << "\033[1;32m"; break;
    case rtcg::PRINT_COLOR::YELLOW   : os << "\033[1;33m"; break;
    case rtcg::PRINT_COLOR::BLUE     : os << "\033[1;34m"; break;
    case rtcg::PRINT_COLOR::MAGENTA  : os << "\033[1;35m"; break;
    case rtcg::PRINT_COLOR::CYAN     : os << "\033[1;36m"; break;
    case rtcg::PRINT_COLOR::WHITE    : os << "\033[1;37m"; break;
    case rtcg::PRINT_COLOR::ENDCOLOR : os << "\033[0m";    break;
    default       : os << "\033[1;37m";
  }
  return os;
}

std::ostream& operator<<(std::ostream& stm, rtcg::Status o)
{
  return stm << rtcg::status::explain(o);
}

bool rtcg::status::isRunning(rtcg::Status s)
{
  switch (s)
  {
  case Status::WAIT_START:
  case Status::WAIT_IN:
  case Status::WAIT_OUT:
    return true;
  default:
    return false;
  }
}

int rtcg::status::exitCode(rtcg::Status s)
{
  switch (s)
  {
  case rtcg::Status::FAIL:
    return 1;
  case rtcg::Status::BUG:
    return 2;
  default:
    return 0;
  }
}

//
// TestCase
//

rtcg::Status rtcg::TestCase::getStatus()
{
  return status_;
}
