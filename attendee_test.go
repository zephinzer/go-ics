package ics

import (
	"testing"
)

func TestAttendeeString(t *testing.T) {
	expected := "ATTENDEE;CN=name"
	attendee := Attendee{Person{
		Name: "name",
	}}
	observed := attendee.String()
	if expected != observed {
		t.Errorf("expected '%s' but got '%s'", expected, observed)
	}
}
