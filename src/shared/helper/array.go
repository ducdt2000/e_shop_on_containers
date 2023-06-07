package array

import (
	"errors"
)

func Find[I interface{}](arr []I, expression func(el I) bool) (*I, error) {
	for _, v := range arr {
		if expression(v) {
			return &v, nil
		}
	}
	return nil, errors.New("Not found")
}

func Filter[I interface{}](arr []I, expression func(el I) bool) ([]I, error) {
	res := []I{}
	for _, v := range arr {
		if expression(v) {
			res = append(res, v)
		}
	}
	return res, nil
}

func FindIndex[I interface{}](arr []I, expression func(el I) bool) (int, error) {
	for i, v := range arr {
		if expression(v) {
			return i, nil
		}
	}
	return -1, errors.New("Not found")
}

func Map[I interface{}, K interface{}](arr []I, expression func(el I) K) ([]K, error) {
	res := []K{}
	for _, v := range arr {
		mapped := expression(v)
		res = append(res, mapped)
	}

	return res, nil
}
