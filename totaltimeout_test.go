package totaltimeout

import "testing"

func TestTimeLeft(t *testing.T) {
	timeout := New(1)
	timeLeft1 := timeout.TimeLeft()
	for timeLeft1 > 0 {
		timeLeft2 := timeout.TimeLeft()
		if timeLeft2 > timeLeft1 {
			t.Error("time left grew")
		}
		timeLeft1 = timeLeft2
	}
}
