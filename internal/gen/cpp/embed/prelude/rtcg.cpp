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
