// Copyright 2019 Alexander Kozhevnikov <mentalisttraceur@gmail.com>
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

// Spread one timeout over many operations.
//
// Correctly and efficiently spreads one timeout over many steps by
// recalculating the time remaining after some amount of waiting has
// already happened, to pass an adjusted timeout to the next step.
//
// Usually idiomatic Go code uses contexts, timers, or deadlines
// instead of timeouts - this helps dealing with code that doesn't.
//
// Example:
//
// Waiting in a "timed loop" for an API with retries (useful for APIs
// that may have transient errors causing immediate error returns):
//
// 	timeout := totaltimeout.New(15 * time.Seconds)
// 	retryDelay := 2 * time.Seconds
// 	for timeLeft := timeout.TimeLeft(); timeLeft > 0 {
// 		client := http.Client{Timeout: timeLeft}
// 		response, err := client.get("https://flaky.api.example.com")
// 		if err == nil && response.StatusCode == 200 {
// 			break
// 		}
// 		if timeout.TimeLeft() < retryDelay {
// 			break
// 		}
// 		time.Sleep(retryDelay)
// 	}
package totaltimeout

import "time"

// Timeout is a single instance of a timeout.
//
// The zero value of a Timeout cannot be used.
type Timeout struct {
	total time.Duration
	start time.Time
	now   func() time.Time
}

// New creates a new Timeout lasting the time.Duration given.
//
// The new Timeout starts immediately, so this should only be
// called at the start of the code the timeout applies to.
func New(timeout time.Duration) Timeout {
	return Timeout{total: timeout, start: time.Now(), now: time.Now}
}

// NewCustomNow is like New but instead of time.Now the new
// Timeout uses the given function to get the current time.
//
// This helps testing and may enable some creative uses.
func NewCustomNow(timeout time.Duration, now func() time.Time) Timeout {
	return Timeout{total: timeout, start: now(), now: now}
}

// TimeLeft gets the time remaining in this Timeout.
func (timeout Timeout) TimeLeft() {
	now := timeout.now()
	elapsed := now.Sub(timeout.start)
	remaining := timeout.total - elapsed
	if remaining > 0 {
		return remaining
	}
	return 0
}
