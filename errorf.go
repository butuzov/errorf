package errorf

import "fmt"

type T struct{}

func (t *T) Errorf(format string, args ...interface{}) {
	println(fmt.Errorf(format, args...).Error())
}
