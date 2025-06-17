package writer

import (
	"fmt"
	"testing"
)

func TestObject(t *testing.T) {
	v := simpleStruct{
		open: true,
		name: "malaka",
	}
	s := v.String()
	fmt.Printf("%s\n", s)
}
