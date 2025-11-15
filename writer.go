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
	w.put(w.fmt(mesg, args...))
}

func (w *Writer) Back(mesg string, args ...any) {
	if w.indent > 0 {
		w.indent--
	}
	w.put(w.fmt(mesg, args...))
}

func (w *Writer) Clear() {
	w.buffer.Reset()
}

func (w *Writer) End(str string, args ...any) {
	w.put(w.fmt(str, args...) + "\n")
}

func (w *Writer) IsEmpty() bool {
	return w.Len() == 0
}

func (w *Writer) Len() int {
	return w.buffer.Len()
}

func (w *Writer) Over(str string, args ...any) {
	str = w.fmt(str, args...)
	if !w.start || str != "" {
		w.put(str + "\n")
	}
	w.indent++
}

func (w *Writer) Sep(chr rune) {
	w.End("# %s", strings.Repeat(string(chr), 78))
}

func (w *Writer) String() string {
	return w.buffer.String()
}

func (w Writer) fmt(str string, args ...any) string {
	if len(args) == 0 {
		return str
	}
	return fmt.Sprintf(str, args...)
}

func (w *Writer) put(str string) {
	prefix := strings.Repeat(w.prefix, w.indent)
	lines := strings.Split(str, "\n")
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
