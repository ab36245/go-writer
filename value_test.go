package writer

import (
	"fmt"
	"testing"
	"time"
)

func TestBytesValue(t *testing.T) {
	v := []byte("I saw a fish going swimming by")
	e := trim(`
	    |30 bytes
        |    0000 49 20 73 61 77 20 61 20 66 69 73 68 20 67 6f 69
        |    0016 6e 67 20 73 77 69 6d 6d 69 6e 67 20 62 79
	`)
	valueTest(t, v, e)
}

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
		hide:     3.14,
		Open:     false,
		malaka:   true,
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
