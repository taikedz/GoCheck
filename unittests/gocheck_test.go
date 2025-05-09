package utest

import (
	"testing"

	"github.com/taikedz/gocheck"
)

func Test_equality(t *testing.T) {
	gocheck.Equal(t, "hello", "hello")
	gocheck.EqualArr(t, []string{"Ami", "Bestie"}, []string{"Ami", "Bestie"})
}
