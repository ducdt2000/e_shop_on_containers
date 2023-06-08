package seed_work

import "reflect"

type ValueObject struct{}

func EqualValueObject(first *ValueObject, second *ValueObject) bool {
	if first == nil || second == nil {
		return false
	}

	return reflect.DeepEqual(first, second)
}
