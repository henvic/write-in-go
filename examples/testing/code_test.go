package clock

import "testing"

func TestTick(t *testing.T) {
	var tick = Tick()

	if tick < 0 {
		t.Errorf("Expected tick to be a positive number, got %v instead", tick)
	}
}
