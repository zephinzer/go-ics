package ics

import (
	"fmt"
	"reflect"
)

type UID struct {
	Value string `json:"uid" ics:"UID"`
}

func (uid *UID) String() string {
	uidType := reflect.TypeOf(UID{})
	valueField, _ := uidType.FieldByName("Value")
	return fmt.Sprintf("%s:%s", valueField.Tag.Get("ics"), uid.Value)
}
