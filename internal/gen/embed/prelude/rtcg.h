//
// rtcg.h - functionality common to rtcg tests.
//
// Copyright (C) 2023 University of York and others
//
// SPDX-License-Identifier: MIT
//

#ifndef RTCG_H_DEFINED
#define RTCG_H_DEFINED

//
// Outcome
//

// An inconclusive, passing, or failing outcome.
enum class Outcome {
    INC,  // Inconclusive outcome.
    PASS, // Passing outcome.
    FAIL  // Failing outcome.
};

// A string representation of an outcome.
const char* outstr(Outcome o);


//
// Status
//

// Statuses of a test-case.
enum class Status
{
  RUNNING    = 0, // Test is still running.
  OFF_SCRIPT = 1, // We saw something that didn't match a transition.
  TIMEOUT    = 2, // We timed out waiting for something to happen.
  FAIL       = 3, // We failed a test.
  BUG        = 4, // Something went internally wrong inside the test.
};

// Returns a (static) description of a status.
const char* explain(Status why);

#endif // RTCG_H_DEFINED
