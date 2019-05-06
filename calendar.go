package ics

const ICalHeader = "BEGIN:VCALENDAR"
const ICalVersion = "VERSION:2.0"
const ICalFooter = "END:VCALENDAR"

type Calendar struct {
	Events []Event `json:"events"`
}

func (calendar *Calendar) Serialize() string {
	return ""
}
