package writer

import (
	"fmt"
	"strings"
)

func New() *Writer {
	return WithPrefix("  ")
}

func WithPrefix(prefix string) *Writer {
	return &Writer{
		prefix: prefix,
		start:  true,
	}
}

type Writer struct {
	buffer strings.Builder
	indent int
	prefix string
	start  bool
}

func (w *Writer) Add(mesg string, args ...any) {
	if len(args) > 0 {
		mesg = fmt.Sprintf(mesg, args...)
	}
	prefix := strings.Repeat(w.prefix, w.indent)
	lines := strings.Split(mesg, "\n")
	for i := range lines {
		line := lines[i]
		if line != "" {
			if w.start {
				w.buffer.WriteString(prefix)
			}
			w.buffer.WriteString(line)
		}
		if i < len(lines)-1 {
			w.buffer.WriteRune('\n')
			w.start = true
		} else {
			w.start = line == ""
		}
	}
}

func (w *Writer) Back(mesg string, args ...any) {
	if w.indent > 0 {
		w.indent--
	}
	if mesg != "" {
		w.Add(mesg, args...)
	}
}

func (w *Writer) Clear() {
	w.buffer.Reset()
}

func (w *Writer) End(mesg string, args ...any) {
	w.Add(mesg+"\n", args...)
}

func (w *Writer) IsEmpty() bool {
	return w.Len() == 0
}

func (w *Writer) Len() int {
	return w.buffer.Len()
}

func (w *Writer) Over(mesg string, args ...any) {
	w.End(mesg, args...)
	w.indent++
}

func (w *Writer) Sep(chr rune) {
	w.End("# %s", strings.Repeat(string(chr), 78))
}

func (w *Writer) String() string {
	return w.buffer.String()
}
