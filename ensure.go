package ensure

import (
	"bytes"
)

type Ensure struct {
	// Used to hold assertion states or results.
	Stack interface{}

	// Accumulates output for writing assertion results, logs, and messages.
	Buffer *bytes.Buffer

	// Determines if assertion checks should stop immediately upon the first failure.
	// If true, assertions stop on the first failure; if false, all assertions run regardless of failures.
	//
	// Defaults to true
	FailFast bool
}
