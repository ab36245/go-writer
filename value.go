package writer

import (
	"reflect"
	"time"
)

func Value() *valueWriter {
	return &valueWriter{
		*New(),
	}
}

type valueWriter struct {
	Writer
}

func (w *valueWriter) Val(value any) *valueWriter {
	w.doVal(value)
	return w
}

func (w *valueWriter) doVal(v any) {
	switch v := v.(type) {
	case bool:
		w.Add("%v", v)
	case int:
		w.Add("%v", v)
	case nil:
		w.Add("nil")
	case string:
		w.Add("%q", v)
	case time.Time:
		w.Add("%q", v.Format(time.RFC3339))
	default:
		r := reflect.ValueOf(v)
		t := r.Type()
		switch t.Kind() {
		case reflect.Map:
			w.doMap(r)
		case reflect.Slice:
			w.doSlice(r)
		default:
			w.Add("??? %v", v)
		}
	}
}

func (w *valueWriter) doMap(r reflect.Value) {
	w.Add("{")
	n := r.Len()
	if n > 0 {
		w.Over("")
		{
			i := r.MapRange()
			for i.Next() {
				w.doVal(i.Key().Interface())
				w.Add(": ")
				w.doVal(i.Value().Interface())
				w.End(",")
			}
		}
		w.Back("")
	}
	w.Add("}")

}

func (w *valueWriter) doSlice(r reflect.Value) {
	w.Add("[")
	n := r.Len()
	if n > 0 {
		w.Over("")
		{
			for i := range n {
				w.doVal(r.Index(i).Interface())
				w.End(",")
			}
		}
		w.Back("")
	}
	w.Add("]")
}
