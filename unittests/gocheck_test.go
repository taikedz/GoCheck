package utest

import (
	"testing"

	"github.com/taikedz/gocheck"
)

func Test_equality(t *testing.T) {
	gocheck.Equal(t, "hello", "hello")
	gocheck.Equal(t, 3, 3.0)

	gocheck.EqualArr(t, []string{"Ami", "Bestie"}, []string{"Ami", "Bestie"})
	gocheck.EqualArr(t, []int{1,2,3}, []int{1.0,2.0,3.0})
}

func Test_difference(t *testing.T) {
	gocheck.Different(t, "hi", "bye")

	gocheck.DifferentArr(t, []int{1,2,3}, []int{1,2})
	gocheck.DifferentArr(t, []int{1,2,3}, []int{1,2,4})
}
