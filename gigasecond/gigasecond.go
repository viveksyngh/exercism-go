package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond add one Giga second to input time and returns it.
func AddGigasecond(t time.Time) time.Time {
	duration := time.Duration(1000000000) * time.Second
	return t.Add(duration)
}
