package interval

import (
	"time"
)

func Run(interval time.Duration, run func()) {
	run()
	for range time.Tick(interval) {
		run()
	}
}

func RunEvery(interval string, run func()) {

	duration, err := time.ParseDuration(interval)
	if err != nil {
		panic(err)
	}

	Run(duration, run)
}

func RunAt(at string, run func()) {

	t, err := time.Parse("3:04pm", at)
	if err != nil {
		panic(err)
	}

	time.Sleep(until(t.Clock()))

	Run(24*time.Hour, run)
}

func until(hour, minute, second int) time.Duration {

	now := time.Now()
	year, month, day := now.Date()

	then := time.Date(
		year,
		month,
		day,
		hour,
		minute,
		second,
		now.Nanosecond(),
		now.Location(),
	)

	if until := then.Sub(now); until < 0 {
		return until + 24*time.Hour
	} else {
		return until
	}
}
