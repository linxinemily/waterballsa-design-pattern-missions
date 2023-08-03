package utils

import "reflect"

func IsSameType(i1 any, i2 any) bool {
	if i1 == nil || i2 == nil {
		return false
	}
	return reflect.TypeOf(i1).String() == reflect.TypeOf(i2).String()
}
