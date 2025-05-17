package writer

import (
	"reflect"
)

type valueWriter struct {
	*Writer
	seen map[uintptr]bool
}

func Value(value any) string {
	w := &valueWriter{
		Writer: New(),
		seen:   make(map[uintptr]bool),
	}
	doValue(w, reflect.ValueOf(value))
	return w.String()
}

func doValue(w *valueWriter, v reflect.Value) {
	t := v.Type()
	switch t.Kind() {
	case reflect.Interface:
		doValue(w, v.Elem())
	case reflect.Map:
		w.Add("{")
		n := v.Len()
		if n > 0 {
			w.Over("")
			{
				i := v.MapRange()
				for i.Next() {
					doValue(w, i.Key())
					w.End(": ")
					doValue(w, i.Value())
					w.End(",")
				}
			}
			w.Back("")
		}
		w.Add("}")
	case reflect.Pointer:
		// Not sure if this is right!
		if v.IsNil() {
			w.Add("nil")
		} else {
			ptr := v.Pointer()
			w.Add("*(0x%08x)", ptr)
			if _, ok := w.seen[ptr]; !ok {
				w.seen[ptr] = true
				w.Add(" ")
				doValue(w, v.Elem())
			}
		}
	case reflect.Slice:
		w.Add("[")
		n := v.Len()
		if n > 0 {
			w.Over("")
			{
				for i := 0; i < n; i++ {
					doValue(w, v.Index(i))
					w.End(",")
				}
			}
			w.Back("")
		}
		w.Add("]")
	case reflect.String:
		w.Add("%q", v)
	case reflect.Struct:
		// w.Add("%s {", v.Type().Name())
		w.Add("%s {", v.Type())
		n := v.NumField()
		if n > 0 {
			w.Over("")
			{
				for i := 0; i < n; i++ {
					w.Add("%s: ", t.Field(i).Name)
					doValue(w, v.Field(i))
					w.End(",")
				}
			}
			w.Back("")
		}
		w.Add("}")
	default:
		w.Add("%v", v)
	}
}
