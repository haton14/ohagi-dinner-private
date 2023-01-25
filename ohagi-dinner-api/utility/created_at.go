package utility

import "time"

func CreatedAt(created *int64, now time.Time) time.Time {
	if created != nil {
		return time.UnixMicro(*created)
	}
	return now
}
