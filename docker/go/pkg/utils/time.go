package utils

import "time"

type Timer interface {
	Now() time.Time
}

type RealTime struct{}

type MockTime struct {
	t time.Time
}

func NewMockTime(t time.Time) MockTime {
	return MockTime{t}
}

func (RealTime) Now() time.Time {
	return time.Now()
}

func (mt MockTime) Now() time.Time {
	return mt.t
}
