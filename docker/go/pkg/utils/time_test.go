package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMockTime(t *testing.T) {
	baseTime := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

	mt := NewMockTime(baseTime)

	assert.Equal(t, baseTime, mt.Now())
}

func TestRealClock(t *testing.T) {
	baseTime := time.Now()
	realTimeNow := RealTime{}.Now()

	assert.Equal(t, realTimeNow.Year(), baseTime.Year())
	assert.Equal(t, realTimeNow.Day(), baseTime.Day())
	assert.Equal(t, realTimeNow.Hour(), baseTime.Hour())
	assert.Equal(t, realTimeNow.Minute(), baseTime.Minute())
}
