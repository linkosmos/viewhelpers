package viewhelpers

import "time"

// Loadtime - load time from give time
func Loadtime(t time.Time) int {
	return int(time.Since(t).Nanoseconds() / 1e6)
}
