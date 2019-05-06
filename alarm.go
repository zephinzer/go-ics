package ics

const ICalAlarmHeader = "BEGIN:VALARM"
const ICalAlarmFooter = "END:VALARM"

type Alarm struct {
	Action      string `json:"action"`
	Description string `json:"description"`
	Trigger     string `json:"trigger"`
	Repeat      int    `json:"repeat"`
}
