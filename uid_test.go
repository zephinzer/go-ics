package ics

import (
	"testing"
)

func TestUID(t *testing.T) {
	expectedValue := "hi"
	expected := "UID:" + expectedValue
	uid := UID{Value: expectedValue}
	observed := uid.String()
	if expected != observed {
		t.Errorf("expected '%s' but got '%s'", expected, observed)
	}
}
