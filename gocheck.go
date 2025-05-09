package gocheck

/*
(C) Tai Kedzierski; Provided under the terms of the MIT License - do what you want with it.

Very simple testing helpers to cut out boilerplate for basic testing.

Either copy it to your project, or add it as a dependency
*/

import (
	"testing"
)

func Equal[V comparable](t *testing.T, exp_value V, got_value V) {
	if exp_value != got_value {
		t.Errorf("Got %v // Exp %v", got_value, exp_value)
	}
}

func EqualArr[V comparable](t *testing.T, exp_value []V, got_value []V) {
	if len(exp_value) != len(got_value) {
		t.Errorf("Got %#v // Exp %#v", got_value, exp_value)
		return
	}

	for i:=0; i<len(exp_value) && i<len(got_value); i++ {
		if exp_value[i] != got_value[i] {
			t.Errorf("Got %#v // Exp %#v", got_value[i], exp_value[i])
			return
		}
	}
}
