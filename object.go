package writer

func Object(name string) *objectWriter {
	return ObjectWithNulls(name, false)
}

func ObjectWithNulls(name string, nulls bool) *objectWriter {
	w := &objectWriter{
		valueWriter: *Value(),
		name:        name,
		nulls:       nulls,
	}
	w.Over("%s{", name)
	return w
}

type objectWriter struct {
	valueWriter
	name  string
	nulls bool
}

func (w *objectWriter) Field(name string, value any) *objectWriter {
	if value != nil || w.nulls {
		w.Add("%s: ", name)
		w.Val(value)
		w.End(",")
	}
	return w
}

func (w *objectWriter) String() string {
	w.Back("}")
	return w.valueWriter.String()
}
