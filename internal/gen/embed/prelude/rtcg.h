//
// rtcg.h - functionality common to rtcg tests
//
// Copyright (C) 2023 University of York and others
//
// SPDX-License-Identifier: MIT
//

#ifndef RTCG_H_DEFINED
#define RTCG_H_DEFINED

namespace rtcg
{
  //
  // Outcome
  //

  // An inconclusive, passing, or failing outcome.
  enum class Outcome {
    UNSET, // Outcome not set.
    INC,   // Inconclusive outcome.
    PASS,  // Passing outcome.
    FAIL   // Failing outcome.
  };

  // A string representation of an outcome.
  const char* outstr(Outcome o);


  //
  // Status
  //

  // Statuses of a test-case.
  enum class Status
  {
    RUNNING,    // Test is still running.
    OFF_SCRIPT, // We saw something that didn't match a transition.
    TIMEOUT,    // We timed out waiting for something to happen.
    FAIL,       // We failed a test.
    BUG,        // Something went internally wrong inside the test.
  };

  // Returns a (static) description of a status.
  const char* explain(Status s);

  // Converts a status into an exit code.
  int exitCode(Status s);
}

#endif // RTCG_H_DEFINED
