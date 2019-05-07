package ics

import (
	"fmt"
	"testing"
)

func TestRecurrenceRuleString(t *testing.T) {
	expectedFrequency := FrequencyYearly
	expectedCount := 2
	expectedInterval := 3
	expected := fmt.Sprintf("RRULE:FREQ=%s;COUNT=%v;INTERVAL=%v", expectedFrequency, expectedCount, expectedInterval)
	rrule := RecurrenceRule{
		Frequency: expectedFrequency,
		Count:     expectedCount,
		Interval:  expectedInterval,
	}
	observed := rrule.String()
	if expected != observed {
		t.Errorf("expected '%s' but got '%s'", expected, observed)
	}
}
