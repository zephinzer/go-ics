package ics

import "fmt"

const ICalHeader = "BEGIN:VCALENDAR"
const ICalFooter = "END:VCALENDAR"

const DefaultCalendarProductID = "zephinzer/go-ics"
const DefaultCalendarCalscale = "GREGORIAN"
const DefaultCalendarMethod = CalendarMethodPublish
const DefaultCalendarVersion = 2.0

const CalendarMethodPublish = "PUBLISH"
const CalendarMethodRequest = "REQUEST"

type Calendar struct {
	Version   float64 `json:"version"`    // https://icalendar.org/iCalendar-RFC-5545/3-7-4-version.html
	Calscale  string  `json:"calscale"`   // https://icalendar.org/iCalendar-RFC-5545/3-7-1-calendar-scale.html
	Method    string  `json:"method"`     // https://icalendar.org/iCalendar-RFC-5545/3-7-2-method.html
	ProductID string  `json:"product_id"` // https://icalendar.org/iCalendar-RFC-5545/3-7-3-product-identifier.html
	Events    []Event `json:"events"`     // https://icalendar.org/iCalendar-RFC-5545/3-6-1-event-component.html
}

func (calendar *Calendar) Serialize() string {
	c := ICalHeader + "\n"

	// Version generation
	if calendar.Version == 0.0 {
		calendar.Version = DefaultCalendarVersion
	}
	c += "VERSION:" + fmt.Sprintf("%f", calendar.Version) + "\n"

	// CalScale generation
	if len(calendar.Calscale) == 0 {
		calendar.Calscale = DefaultCalendarCalscale
	}
	c += "CALSCALE:" + calendar.Calscale + "\n"

	// Method generation
	if len(calendar.Method) == 0 {
		calendar.Method = DefaultCalendarMethod
	}
	c += "METHOD:" + calendar.Method + "\n"

	// ProductID generation
	if len(calendar.ProductID) == 0 {
		calendar.ProductID = DefaultCalendarProductID
	}
	c += "PRODID:" + calendar.ProductID + "\n"

	c += ICalFooter + "\n"
	return c
}
