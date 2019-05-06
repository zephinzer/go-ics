package ics

// Attendee implements the attendee object
// See: https://icalendar.org/iCalendar-RFC-5545/3-8-4-1-attendee.html
type Attendee struct {
	person Person
}

func (attendee Attendee) String() string {
	return "ATTENDEE;" + attendee.person.String()
}
