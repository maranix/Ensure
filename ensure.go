package ensure

import (
	"bytes"
)

type AssertionStatus int8

const (
    AssertionSuccess AssertionStatus = iota
    AssertionFailure 
)

func (as AssertionStatus) String() string {
    switch as {
    case AssertionSuccess:
        return "Success"
    case AssertionFailure:
        return "Failure"
    }

    panic("Invalid AssertionStatus was obtained\n\nThis should never happen in any case. Please file an issue on GitHub repo along with reproduction steps")
}


type Ensure struct {
    // Used to hold assertion states or results. 
    stack interface{}

    // Accumulates output for writing assertion results, logs, and messages.
    buffer *bytes.Buffer

    // Determines if assertion checks should stop immediately upon the first failure.
    // If true, assertions stop on the first failure; if false, all assertions run regardless of failures.
    //
    // Defaults to true
    failFast bool
}
