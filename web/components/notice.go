package components

import "time"

// Notice is a time-boxed site-wide announcement.
type Notice struct {
	From    time.Time // inclusive
	Until   time.Time // exclusive
	Message string
}

// Active reports whether the notice should be shown at t.
func (n Notice) Active(t time.Time) bool {
	return !t.Before(n.From) && t.Before(n.Until)
}

// BreakNotice is the seasonal "on a break" banner. The Nuxt layout showed it
// while `getMonth() <= 7 && getFullYear() == 2026`, i.e. Jan 1 – Aug 31 2026.
var BreakNotice = Notice{
	From:    time.Date(2026, time.January, 1, 0, 0, 0, 0, time.Local),
	Until:   time.Date(2026, time.September, 1, 0, 0, 0, 0, time.Local),
	Message: "We will be on break until the 01/09/2026. Please be aware that the orders you made will be shipped after 01/09/2026",
}
