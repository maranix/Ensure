package ensure

import (
	"bytes"
	"fmt"
)

type AssertStatus int8

const (
	AssertSuccess AssertStatus = iota
	AssertFailure
	AssertUnknown
)

func (as AssertStatus) String() string {
	switch as {
	case AssertSuccess:
		return "Success"
	case AssertFailure:
		return "Failure"
	default:
		return "Unknown"
	}
}

func (as AssertStatus) ColoredString() string {
	switch as {
	case AssertSuccess:
		return colorizeString(green, "Success")
	case AssertFailure:
		return colorizeString(red, "Failure")
	default:
		return colorizeString(yellow, "Unknown")
	}
}

type AssertOperation string

const (
	Equal      AssertOperation = "Equal"
	NotEqual                   = "NotEqual"
	IsNil                      = "IsNil"
	NotNil                     = "NotNil"
	IsEmpty                    = "IsEmpty"
	NotEmpty                   = "NotEmpty"
	IsType                     = "IsType"
	InRange                    = "InRange"
	NotInRange                 = "NotInRange"
)

type AssertKind int8

const (
	EqualityCheck  AssertKind = iota // Checks if two values are equal or not
	NilCheck                         // Checks if a value is nil or not
	EmptinessCheck                   // Checks if a value is empty or not
	TypeCheck                        // Checks the type of a value
	RangeCheck                       // Checks if a value is within a range
)

func (ak AssertKind) String() string {
	switch ak {
	case EqualityCheck:
		return "Equality Check"
	case NilCheck:
		return "Nil Check"
	case EmptinessCheck:
		return "Emptiness Check"
	case TypeCheck:
		return "Type Check"
	case RangeCheck:
		return "Range Check"
	}

	panic(
		"Invalid AssertKind received" +
			"\n" +
			"This should never happen in any circumstance." +
			"\n" +
			"Please contact the author and file an issue on GitHub with either of the following information:" +
			"\n\t1. Minimal sample application/usecase which causes this panic." +
			"\n\t2. Steps to reproduce this panic in an isolated environment.")
}

type Assert struct {
	Id      int
	Op      AssertOperation
	Kind    AssertKind
	Have    any
	Want    any
	Status  AssertStatus
	Message string
}

func NewAssert(id int, op AssertOperation, kind AssertKind, have any, want any, message string) *Assert {
	return &Assert{
		Id:      id,
		Op:      op,
		Kind:    kind,
		Have:    have,
		Want:    want,
		Status:  AssertUnknown,
		Message: message,
	}
}

func (a *Assert) String() string {
	b := new(bytes.Buffer)

	if a.Message != "" {
		b.WriteString(a.Message)
		b.WriteString("\n")
	}

	b.WriteString(fmt.Sprintf("Status: %s, ", a.Status.ColoredString()))
	b.WriteString(fmt.Sprintf("Operation: %s, ", a.Op))
	b.WriteString(fmt.Sprintf("Kind: %s, ", a.Kind.String()))
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("Have: %v, ", a.Have))
	b.WriteString(fmt.Sprintf("Want: %v, ", a.Want))

	return b.String()
}
