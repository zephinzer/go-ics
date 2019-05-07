package ics

import (
	"fmt"
	"math"
	"strings"
	"time"
)

const ICalAlarmHeader = "BEGIN:VALARM"
const ICalAlarmFooter = "END:VALARM"

const DefaultAlarmAction = AlarmActionDisplay

const (
	AlarmActionAudio     = "AUDIO"
	AlarmActionDisplay   = "DISPLAY"
	AlarmActionEmail     = "EMAIL"
	AlarmActionProcedure = "PROCEDURE"
)

var (
	AlarmActions = []string{
		AlarmActionAudio,
		AlarmActionDisplay,
		AlarmActionEmail,
		AlarmActionProcedure,
	}
)

type Alarm struct {
	Action      string        `json:"action"`
	Description string        `json:"description"`
	Trigger     time.Duration `json:"trigger"`
	Repeat      int           `json:"repeat"`
}

func (alarm *Alarm) Serialize() string {
	a := ICalAlarmHeader + "\n"

	// Action
	if len(alarm.Action) == 0 {
		alarm.Action = DefaultAlarmAction
	}
	a += "ACTION:" + strings.ToUpper(alarm.Action) + "\n"

	// Trigger
	if alarm.Trigger != time.Duration(0) {
		d := alarm.Trigger.Round(time.Minute).Minutes()
		minutes := math.Mod(d, 60)
		hours := (d - minutes) / 60
		triggerString := fmt.Sprintf("TRIGGER:-PT%vH%vM", hours, minutes)
		a += triggerString + "\n"
	}

	// Description
	if len(alarm.Description) > 0 {
		a += "DESCRIPTION:" + alarm.Description + "\n"
	}

	// Repeat
	if alarm.Repeat > 0 {
		a += fmt.Sprintf("REPEAT:%v", alarm.Repeat) + "\n"
	}

	a += ICalAlarmFooter
	return a
}
