package util

// ExecIfErrorFunc function type used in ExecIfError method
type ExecIfErrorFunc func(e error)

// ExecIfNotErrorFunc function type used in ExecIfNotErrorFunc method
type ExecIfNotErrorFunc func()

// IsError eturns true if the error value is not nil, false
// otherwise.
func IsError(e error) bool {
	if e != nil {
		return true
	}

	return false
}

// IsNotError returns true if the error value is nil, false
// otherwise.
func IsNotError(e error) bool {
	return !IsError(e)
}

// ExecIfError calls the passed ExecIfErrorFunc function if
// the error value is not nil.
func ExecIfError(e error, f ExecIfErrorFunc) {
	if IsError(e) {
		f(e)
	}
}

// ExecIfNotError calls the passed ExecIfNotErrorFunc function if
// the error value is not nil.
func ExecIfNotError(e error, f ExecIfNotErrorFunc) {
	if IsNotError(e) {
		f()
	}
}
