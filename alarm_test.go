package ics

import (
	"testing"
	"time"
)

func TestAlarmSerialize(t *testing.T) {
	expectedAction := AlarmActionDisplay
	expectedDescription := "lorem ipsum"
	expectedTrigger := time.Hour*3 + time.Minute*23 + time.Second*58
	expected := ICalAlarmHeader + "\n" +
		"ACTION:" + expectedAction + "\n" +
		"DESCRIPTION:" + expectedDescription + "\n" +
		"TRIGGER:-PT3H24M" + "\n" +
		ICalAlarmFooter
	alarm := Alarm{
		Action:      expectedAction,
		Description: expectedDescription,
		Trigger:     expectedTrigger,
	}
	observed := alarm.Serialize()
	if expected != observed {
		t.Errorf("expected '%s' but got '%s'", expected, observed)
	}
}
