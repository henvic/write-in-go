package clock

var current = -1

func Tick() int {
	current += 1
	return current
}
