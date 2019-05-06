package ics

import (
	"fmt"

	"github.com/google/uuid"
)

const ICalEventHeader = "BEGIN:VEVENT"
const ICalEventFooter = "END:VEVENT"

// Event (https://icalendar.org/iCalendar-RFC-5545/3-6-1-event-component.html)
type Event struct {
	UID            string    `json:"uid"`
	Title          string    `json:"title"`
	Summary        string    `json:"summary"`     // https://icalendar.org/iCalendar-RFC-5545/3-8-1-12-summary.html
	Description    string    `json:"description"` // https://icalendar.org/iCalendar-RFC-5545/3-8-1-5-description.html
	URL            string    `json:"url"`
	Categories     []string  `json:"categories"`
	Status         string    `json:"status"`
	Organizer      Person    `json:"organizer"`
	Location       Location  `json:"geo"`
	Timestamp      Timestamp `json:"timestamp"`
	Sequence       int       `json:"sequence"`
	RecurrenceRule string    `json:"recurrence_rule"`
}

func (event *Event) Serialize() string {
	e := ICalEventHeader + "\n"

	// UID generation
	if len(event.UID) == 0 {
		event.UID = uuid.New().String()
	}
	e += "UID:" + event.UID + "\n"

	// Title
	if len(event.Title) > 0 {
		e += "TITLE:" + event.Title
	}

	// Summary
	if len(event.Summary) > 0 {
		e += "SUMMARY:" + event.Summary
	}

	// Description
	if len(event.Description) > 0 {
		e += "DESCRIPTION:" + event.Description
	}

	// Summary
	if len(event.Summary) > 0 {
		e += "SUMMARY:" + event.Summary
	}

	// Sequence
	e += "SEQUENCE:" + fmt.Sprintf("%v", event.Sequence) + "\n"

	e += ICalEventFooter + "\n"
	return e
}
