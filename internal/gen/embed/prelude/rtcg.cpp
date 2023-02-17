//
// rtcg.cpp - functionality common to rtcg tests.
//
// Copyright (C) 2023 University of York and others
//
// SPDX-License-Identifier: MIT
//

#include "rtcg.h"

//
// Outcome
//

const char* outstr(Outcome o)
{
  switch (o) {
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

const char* explain(Status why)
{
  switch (why) {
  case Status::RUNNING:
    return "running";
  case Status::OFF_SCRIPT:
    return "event happened that took the test off-script";
  case Status::TIMEOUT:
    return "no new events happening within the allotted timeframe";
  case Status::FAIL:
    return "an unwanted behaviour occurring";
  case Status::BUG:
    return "internal error";
  default:
    return "unknown (this should not occur)";
  }
}