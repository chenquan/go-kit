package log

type Field interface {
	Key() string
	Value() any
}

type field struct {
	key   string
	value any
}

func (f field) Key() string {
	return f.key
}

func (f field) Value() any {
	return f.value
}

func FieldKv(key string, value any) Field {
	return field{
		key:   key,
		value: value,
	}
}
