Spread one timeout over many operations.

Correctly and efficiently spreads one timeout over many steps by
recalculating the time remaining after some amount of waiting has
already happened, to pass an adjusted timeout to the next step.

Usually idiomatic Go code uses contexts, timers, or deadlines
instead of timeouts - this helps dealing with code that doesn't.
