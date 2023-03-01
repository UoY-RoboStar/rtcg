//
// rtcg.cpp - functionality common to rtcg tests
//
// Copyright (C) 2023 University of York and others
//
// SPDX-License-Identifier: MIT
//

#include "rtcg.h"

//
// Outcome
//

const char* rtcg::outstr(rtcg::Outcome o)
{
  switch (o) {
  case Outcome::UNSET:
    return "unset";
  case Outcome::INC:
    return "inconclusive";
  case Outcome::PASS:
    return "passed";
  case Outcome::FAIL:
    return "failed";
  }
  return "???";
}


//
// Status
//

const char* rtcg::explain(Status s)
{
  switch (s) {
  case Status::RUNNING:
    return "still running";
  case Status::OFF_SCRIPT:
    return "an event happened that took the test off-script";
  case Status::TIMEOUT:
    return "the test ran out of time";
  case Status::FAIL:
    return "an unwanted behaviour occurred";
  case Status::BUG:
    return "an internal error occurred";
  default:
    return "unknown (this should not occur)";
  }
}

int rtcg::exitCode(Status s)
{
  switch (s) {
  case Status::FAIL:
    return 1;
  case Status::BUG:
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
