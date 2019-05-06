package ics

import "time"

type Timestamp struct {
	Created  time.Time `json:"created"`
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
	Duration time.Time `json:"duration"`
}
