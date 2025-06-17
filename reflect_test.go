package writer

import (
	"testing"
)

func TestReflectSlice(t *testing.T) {
	v := []int{2, 3}
	a := Reflect(v)
	e := "[\n  2,\n  3,\n]"
	if a != e {
		t.Fatalf("expected %s, actual %s", e, a)
	}
}

func TestReflectString(t *testing.T) {
	v := "malaka"
	a := Reflect(v)
	e := `"malaka"`
	if a != e {
		t.Fatalf("expected %s, actual %s", e, a)
	}
}

func TestReflectStruct(t *testing.T) {
	v := simpleStruct{
		open: false,
		name: "malaka",
	}
	a := Reflect(v)
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
