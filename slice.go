package gu
// ForEach applies the function f to each element of the slice s.
func ForEach[T any](s []T, f func(T)) {
	for i := 0; i < len(s); i++ {
		f(s[i])
	}
}

// Map returns a new slice containing the results of applying the function f to each element of the slice s.
func Map[T, U any](s []T, f func(T) U) []U {
	result := make([]U, len(s))
	for i := 0; i < len(s); i++ {
		result[i] = f(s[i])
	}
	return result
}

func Index[T comparable](s []T, v T, reverse ...bool) int {
	if len(reverse) > 0 && reverse[0] {
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] == v {
				return i
			}
		}
		return -1
	}
	for i := 0; i < len(s); i++ {
		if s[i] == v {
			return i
		}
	}
	return -1
}

func Contains[T comparable](s []T, v T) bool {
	return Index(s, v) != -1
}

func All[T any](s []T, f func(T) bool) bool {
	for i := 0; i < len(s); i++ {
		if !f(s[i]) {
			return false
		}
	}
	return true
}

func Any[T any](s []T, f func(T) bool) bool {
	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			return true
		}
	}
	return false
}

func Count[T any](s []T, f func(T) bool) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			count++
		}
	}
	return count
}

func Where[T any](s []T, f func(T) bool) []T {
	var result []T
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func FirstWhere[T any](s []T, f func(T) bool) (T, bool) {
    for _, v := range s {
        if f(v) {
            return v, true
        }
    }
    var zero T
    return zero, false
}

func LastWhere[T any](s []T, f func(T) bool) (T, bool) {
	for i := len(s) - 1; i >= 0; i-- {
		if f(s[i]) {
			return s[i], true
		}
	}
	var zero T
	return zero, false
}

func Reduce[T any, R any](s []T, f func(agg R, item T, index int) R, initial R, fromRight ...bool) R {
	if len(fromRight) > 0 && fromRight[0] {
		for i := len(s) - 1; i >= 0; i-- {
			initial = f(initial, s[i], i)
		}
		return initial
	}
	for i := 0; i < len(s); i++ {
		initial = f(initial, s[i], i)
	}

	return initial
}

func While[T any, R any](s []T, f func(T) (R, bool)) []R {
	var result []R
	for i := 0; i < len(s); i++ {
		v := s[i]
		r, ok := f(v)
		if !ok {
			break
		}
		result = append(result, r)
	}
	return result
}
