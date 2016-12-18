package String

import "strings"

type self *String
type String struct {
	source []byte
}

func New(s interface{}) self {
	var b []byte

	switch s.(type) {
	case string:
		b = []byte(s)
	case []byte:
		b = s
	default:
		panic(`Unknow type`)
	}

	return &String{
		source: b,
	}
}

func (s self) String() string {
	return string(s.source)
}

func (s self) Contains(sub interface{}) {
	return strings.Contains(s.String(), New(sub).String())
}

func (s self) Split(seq interface{}) []self {
	return strings.Split(s.String(), New(seq).String())
}
