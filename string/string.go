package string

import (
	"fmt"
	"strings"
)

type String []byte

func New(s interface{}) String {
	switch s.(type) {
	case *string:
		return New([]byte(*s.(*string)))
	case string:
		return New([]byte(s.(string)))
	case []byte:
		return String(s.([]byte))
	case String:
		return s.(String)
	default:
		return New(fmt.Sprint(s))
	}
}

func (s String) String() string {
	return string(s)
}

func (s String) bytes() []byte {
	return s
}

func (s String) Equal(sub interface{}) bool {
	return s.String() == New(sub).String()
}

func (s String) Contains(sub interface{}) bool {
	return strings.Contains(s.String(), New(sub).String())
}

func (s String) Index(sub interface{}) int {
	return strings.Index(s.String(), New(sub).String())
}

func (s String) Split(seq interface{}) []string {
	return strings.Split(s.String(), New(seq).String())
}

func (s String) HasPrefix(seq interface{}) bool {
	return strings.HasPrefix(s.String(), New(seq).String())
}

func (s String) HasSuffix(seq interface{}) bool {
	return strings.HasSuffix(s.String(), New(seq).String())
}

func (s String) Trim(seq interface{}) String {
	return s.Ltrim(seq).Rtrim(seq)
}

func (s String) Ltrim(seq interface{}) String {
	return s
}

func (s String) Rtrim(seq interface{}) String {
	return s
}
