// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2017 Datadog, Inc.

package retry

import "time"

// Status is returned by Retryer object to inform user classes
type Status int

const (
	// NeedSetup is the default value: SetupRetryer must be called
	NeedSetup Status = iota // Default zero value
	// Idle means the Retryer is ready for Try to be called
	Idle
	// OK means the object is available
	OK
	// FailWillRetry informs users the object is not avalable yet,
	// but they should retry later
	FailWillRetry
	// PermaFail informs the user the object will not be available.
	PermaFail
)

// Strategy sets how the Retryer should handle failure
type Strategy int

const (
	// OneTry is the default value: only try one, then permafail
	OneTry Strategy = iota // Default zero value
	// RetryCount sets the Retryer to try a fixed number of times
	RetryCount
	// RetryDuration sets the Retryer to try for a fixed duration
	// RetryDuration // FIXME: implement
)

// Config contains all the required parameters for Retryer
type Config struct {
	Name          string
	AttemptMethod func() error
	Strategy      Strategy
	RetryCount    int
	RetryDelay    time.Duration
}
