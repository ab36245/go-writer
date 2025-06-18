package writer

import (
	"fmt"
	"reflect"
	"time"
)

func Value(value any) string {
	w := WithPrefix("    ")
	doVal(w, value)
	return w.String()
}

func doVal(w *Writer, v any) {
	switch v := v.(type) {
	case nil:
		w.Add("nil")
	case bool:
		w.Add("%v", v)
	case float64:
		w.Add("%v", v)
	case int:
		w.Add("%v", v)
	case string:
		w.Add("%q", v)
	case time.Time:
		w.Add("%s", v.Format(time.RFC3339))
	default:
		r := reflect.ValueOf(v)
		switch r.Type().Kind() {
		case reflect.Map:
			doMap(w, r)
		case reflect.Slice:
			doSlice(w, r)
		case reflect.Struct:
			doStruct(w, r)
		default:
			w.Add("??? %v", v)
		}
	}
}

func doMap(w *Writer, r reflect.Value) {
	w.Add("{")
	n := r.Len()
	if n > 0 {
		w.Over("")
		{
			i := r.MapRange()
			for i.Next() {
				doVal(w, i.Key().Interface())
				w.Add(": ")
				doVal(w, i.Value().Interface())
				w.End(",")
			}
		}
		w.Back("")
	}
	w.Add("}")

}

func doSlice(w *Writer, r reflect.Value) {
	w.Add("[")
	n := r.Len()
	if n > 0 {
		w.Over("")
		{
			for i := range n {
				doVal(w, r.Index(i).Interface())
				w.End(",")
			}
		}
		w.Back("")
	}
	w.Add("]")
}

func doStruct(w *Writer, v reflect.Value) {
	vt := v.Type()
	w.Add("%s{", vt.Name())

	var fs []reflect.Value
	for i := range v.NumField() {
		f := v.Field(i)
		if f.CanInterface() {
			// This is only true for exported fields
			fs = append(fs, f)
		}
	}

	if len(fs) > 0 {
		w.Over("")
		{
			// Do go-like justifying of field names
			max := 0
			for i := range fs {
				ft := vt.Field(i)
				if max < len(ft.Name) {
					max = len(ft.Name)
				}
			}
			for i, f := range fs {
				ft := vt.Field(i)
				name := fmt.Sprintf("%s:", ft.Name)
				w.Add("%-*s ", max+1, name)
				doVal(w, f.Interface())
				w.End(",")
			}
		}
		w.Back("")
	}
	w.Add("}")
}
