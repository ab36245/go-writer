package writer

import (
	"fmt"
	"testing"
	"time"
)

func TestValueInt(t *testing.T) {
	v := 42
	e := "42"
	valueTest(t, v, e)
}

func TestValueMap(t *testing.T) {
	v := map[int]string{
		1:  "one",
		2:  "two",
		99: "ninety-nine",
	}
	e := trim(`
	    |{
        |    1: "one",
        |    2: "two",
        |    99: "ninety-nine",
        |}
	`)
	valueTest(t, v, e)
}

func TestValueNil(t *testing.T) {
	valueTest(t, nil, "nil")
}

func TestValueSlice(t *testing.T) {
	v := []float64{3.14159, 2.71828}
	e := trim(`)
	    |[
        |    3.14159,
        |    2.71828,
        |]
	`)
	valueTest(t, v, e)
}

func TestValueString(t *testing.T) {
	v := "malaka was here"
	e := fmt.Sprintf("%q", v)
	valueTest(t, v, e)
}

func TestValueStruct(t *testing.T) {
	v := simpleStruct{
		LongName: "malaka",
		Open:     false,
	}
	e := trim(`
		|simpleStruct{
		|    Open:     false,
		|    LongName: "malaka",
		|}
	`)
	valueTest(t, v, e)
}

func TestValueTime(t *testing.T) {
	v, _ := time.Parse(time.DateOnly, "1997-08-28")
	e := "1997-08-28T00:00:00Z"
	valueTest(t, v, e)
}

func valueTest(t *testing.T, v any, e string) {
	s := Value(v)
	if s != e {
		t.Fatalf("expected %s, actual %s", e, s)
	}
}
