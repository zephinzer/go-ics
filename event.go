package ics

import (
	"fmt"
	"reflect"

	"github.com/google/uuid"
)

const ICalEventHeader = "BEGIN:VEVENT"
const ICalEventFooter = "END:VEVENT"

// Event (https://icalendar.org/iCalendar-RFC-5545/3-6-1-event-component.html)
type Event struct {
	UID            *UID            `json:"uid"`
	Title          string          `json:"title" ics:"TITLE:"`
	Summary        string          `json:"summary" ics:"SUMMARY:"`         // https://icalendar.org/iCalendar-RFC-5545/3-8-1-12-summary.html
	Description    string          `json:"description" ics:"DESCRIPTION:"` // https://icalendar.org/iCalendar-RFC-5545/3-8-1-5-description.html
	URL            string          `json:"url" ics:"URL"`
	Categories     []string        `json:"categories"`
	Status         string          `json:"status"`
	Sequence       int             `json:"sequence" ics:"SEQUENCE"`
	Organizer      Person          `json:"organizer"`
	Location       Location        `json:"geo"`
	Timestamp      Timestamp       `json:"timestamp"`
	RecurrenceRule *RecurrenceRule `json:"recurrence_rule"`
}

func (event *Event) Serialize() string {
	eventType := reflect.TypeOf(Event{})
	e := ICalEventHeader + "\n"

	// UID generation
	if event.UID == nil {
		event.UID = &UID{Value: uuid.New().String()}
	}
	e += event.UID.String() + "\n"

	// Title
	if len(event.Title) > 0 {
		titleField, _ := eventType.FieldByName("Title")
		e += fmt.Sprintf("%s:%s\n", titleField.Tag.Get("ics"), event.Title)
	}

	// Summary
	if len(event.Summary) > 0 {
		summaryField, _ := eventType.FieldByName("Summary")
		e += fmt.Sprintf("%s:%s\n", summaryField.Tag.Get("ics"), event.Summary)
	}

	// Description
	if len(event.Description) > 0 {
		descriptionField, _ := eventType.FieldByName("Description")
		e += fmt.Sprintf("%s:%s\n", descriptionField.Tag.Get("ics"), event.Description)
	}

	// Sequence
	sequenceField, _ := eventType.FieldByName("Sequence")
	e += fmt.Sprintf("%s:%v", sequenceField.Tag.Get("ics"), event.Sequence) + "\n"

	// RecurrenceRule
	if event.RecurrenceRule != nil {
		e += event.RecurrenceRule.String() + "\n"
	}

	e += ICalEventFooter + "\n"
	return e
}
