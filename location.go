package ics

import (
	"fmt"
	"strings"
)

type Location struct {
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (location Location) String() string {
	l := ""
	if len(location.Address) > 0 {
		l += fmt.Sprintf("LOCATION:%s\n", location.Address)
	}
	if location.Latitude > 0.0 && location.Longitude > 0.0 {
		l += fmt.Sprintf("GEO:%f;%f\n", location.Latitude, location.Longitude)
	}
	return strings.Trim(l, "\n ")
}
