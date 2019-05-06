package ics

import (
	"fmt"
	"testing"
)

func TestLocationString_Address(t *testing.T) {
	expectedAddress := "1 Address #08-01, City 123123"
	expected := "LOCATION:" + expectedAddress
	location := Location{Address: expectedAddress}
	observed := location.String()
	if expected != observed {
		t.Errorf("expected '%s' but got '%s'", expected, observed)
	}
}

func TestLocationString_AddressLatitude(t *testing.T) {
	expectedAddress := "1 Address #08-01, City 123123"
	expectedLatitude := 1.123
	expected := "LOCATION:" + expectedAddress
	location := Location{Address: expectedAddress, Latitude: expectedLatitude}
	observed := location.String()
	if expected != observed {
		t.Errorf("expected '%s' but got '%s'", expected, observed)
	}
}

func TestLocationString_AddressLongitude(t *testing.T) {
	expectedAddress := "1 Address #08-01, City 123123"
	expectedLongitude := 1.123
	expected := "LOCATION:" + expectedAddress
	location := Location{Address: expectedAddress, Longitude: expectedLongitude}
	observed := location.String()
	if expected != observed {
		t.Errorf("expected '%s' but got '%s'", expected, observed)
	}
}

func TestLocationString_AddressLatitudeLongitude(t *testing.T) {
	expectedAddress := "1 Address #08-01, City 123123"
	expectedLatitude := 2.234
	expectedLongitude := 1.123
	expected := "LOCATION:" + expectedAddress + "\n" + "GEO:" + fmt.Sprintf("%f;%f", expectedLatitude, expectedLongitude)
	location := Location{Address: expectedAddress, Latitude: expectedLatitude, Longitude: expectedLongitude}
	observed := location.String()
	if expected != observed {
		t.Errorf("expected '%s' but got '%s'", expected, observed)
	}
}
