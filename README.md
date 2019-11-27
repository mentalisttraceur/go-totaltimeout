Spread one timeout over many operations.

Correctly and efficiently spreads one timeout over many steps by
recalculating the time remaining after some amount of waiting has
already happened, to pass an adjusted timeout to the next step.

Usually idiomatic Go code uses contexts, timers, or deadlines
instead of timeouts - this helps dealing with code that doesn't.

Example:

Waiting in a "timed loop" for an API with retries (useful for APIs
that may have transient errors causing immediate error returns):

	timeout := totaltimeout.New(15 * time.Seconds)
	retryDelay := 2 * time.Seconds
	for timeLeft := timeout.TimeLeft(); timeLeft > 0 {
		client := http.Client{Timeout: timeLeft}
		response, err := client.get("https://flaky.api.example.com")
		if err == nil && response.StatusCode == 200 {
			break
		}
		if timeout.TimeLeft() < retryDelay {
			break
		}
		time.Sleep(retryDelay)
	}
