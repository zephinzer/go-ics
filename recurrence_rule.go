package ics

import (
	"fmt"
	"strings"
)

const (
	FrequencyDaily   = "DAILY"
	FrequencyMonthly = "MONTHLY"
	FrequencyYearly  = "YEARLY"
)

type RecurrenceRule struct {
	Frequency string `json:"frequency"`
	Count     int    `json:"count"`
	Interval  int    `json:"interval"`
}

func (rrule *RecurrenceRule) String() string {
	r := "RRULE:"

	// Frequency
	if len(rrule.Frequency) > 0 {
		r += "FREQ=" + strings.ToUpper(rrule.Frequency)
	}

	// Count
	if rrule.Count > 0 {
		r += fmt.Sprintf(";COUNT=%v", rrule.Count)
	}

	// Interval
	if rrule.Interval > 0 {
		r += fmt.Sprintf(";INTERVAL=%v", rrule.Interval)
	}

	return r
}
