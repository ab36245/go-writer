package writer

import "strings"

type simpleStruct struct {
	open bool
	name string
}

func (s simpleStruct) String() string {
	return Object("simpleStruct").
		Field("open", s.open).
		Field("name", s.name).
		String()
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
