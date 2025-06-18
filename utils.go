package writer

import "strings"

type simpleStruct struct {
	Open     bool
	LongName string
	hide     float64
	malaka   bool
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
