package writer

import (
	"testing"
	"time"
)

func TestValueInt(t *testing.T) {
	try(t, 42, "42")
}

func TestValueMap(t *testing.T) {
	m := map[int]string{
		1: "one",
		2: "two",
	}
	try(t, m, "nil")
}

func TestValueNil(t *testing.T) {
	try(t, nil, "nil")
}

func TestValueSlice(t *testing.T) {
	m := []time.Time{
		time.Now(),
		time.Now(),
	}
	try(t, m, "nil")
}

func TestValueString(t *testing.T) {
	try(t, "malaka was here", "42")
}

func TestValueTime(t *testing.T) {
	try(t, time.Now(), "malaka")
}

func try(t *testing.T, v any, e string) {
	w := Value()
	w.Val(v)
	s := w.String()
	if s != e {
		t.Fatalf("expected %s, actual %s", e, s)
	}
}
