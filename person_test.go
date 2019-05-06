package ics

import (
	"testing"
)

func TestPersonString_Name(t *testing.T) {
	expectedName := "Joseph Goh"
	expected := "CN=" + expectedName
	person := Person{Name: expectedName}
	observed := person.String()
	if expected != observed {
		t.Errorf("expected '%s' but got '%s'", expected, observed)
	}
}

func TestPersonString_NameDir(t *testing.T) {
	expectedName := "Joseph Goh"
	expectedDir := "https://github.com/zephinzer"
	expected := "CN=" + expectedName + ";DIR=\"" + expectedDir + "\""
	person := Person{Name: expectedName, Dir: expectedDir}
	observed := person.String()
	if expected != observed {
		t.Errorf("expected '%s' but got '%s'", expected, observed)
	}
}

func TestPersonString_NameEmail(t *testing.T) {
	expectedName := "Joseph Goh"
	expectedEmail := "email@domain.com"
	expected := "CN=" + expectedName + ":mailto:" + expectedEmail
	person := Person{Name: expectedName, Email: expectedEmail}
	observed := person.String()
	if expected != observed {
		t.Errorf("expected '%s' but got '%s'", expected, observed)
	}
}

func TestPersonString_NameEmailDir(t *testing.T) {
	expectedName := "Joseph Goh"
	expectedEmail := "email@domain.com"
	expectedDir := "https://github.com/zephinzer"
	expected := "CN=" + expectedName + ";DIR=\"" + expectedDir + "\":mailto:" + expectedEmail
	person := Person{Name: expectedName, Email: expectedEmail, Dir: expectedDir}
	observed := person.String()
	if expected != observed {
		t.Errorf("expected '%s' but got '%s'", expected, observed)
	}
}
