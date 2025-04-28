package writer

import (
	"strings"
	"testing"
)

func TestSliceValue(t *testing.T) {
	v := []int{2, 3}
	a := Value(v)
	e := "[\n  2,\n  3,\n]"
	if a != e {
		t.Fatalf("expected %s, actual %s", e, a)
	}
}

func TestStringValue(t *testing.T) {
	v := "malaka"
	a := Value(v)
	e := `"malaka"`
	if a != e {
		t.Fatalf("expected %s, actual %s", e, a)
	}
}

type simpleStruct struct {
	open bool
	name string
}

func TestStructValue(t *testing.T) {
	v := simpleStruct{
		open: false,
		name: "malaka",
	}
	a := Value(v)
	e := trim(`
		|simpleStruct {
		|  open: false,
		|  name: "malaka",
		|}
	`)
	if a != e {
		t.Fatalf("expected %s, actual %s", e, a)
	}
}

func trim(s string) string {
	t := ""
	for l := range strings.SplitSeq(s, "\n") {
		m := strings.TrimSpace(l)
		if n, found := strings.CutPrefix(m, "|"); found {
			if t != "" {
				t += "\n"
			}
			t += n
		}
	}
	return t
}
