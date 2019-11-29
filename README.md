Spread one timeout over many operations.

Correctly and efficiently spreads one timeout over many steps by
recalculating the time remaining after some amount of waiting has
already happened, to pass an adjusted timeout to the next step.

Usually idiomatic Go code uses contexts, timers, or deadlines
instead of timeouts - this helps dealing with code that doesn't.

Examples:

Waiting in a "timed loop" for an API with retries (useful
for unreliable APIs that may either hang or need retries):

	timeout := totaltimeout.New(15 * time.Seconds)
	for timeLeft := timeout.TimeLeft(); timeLeft > 0 {
		client := http.Client{Timeout: timeLeft}
		response, err := client.get("https://flaky.api.example.com")
		if err == nil && response.StatusCode == 200 {
			break
		}
	}
	// handle success or failure here, once all retries are done

Same as above, but with a wait between retries:

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
