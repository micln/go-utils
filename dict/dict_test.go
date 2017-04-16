package dict

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"testing"
)

//	test Getter/Setter
func TestNewDict(t *testing.T) {
	d := NewDict()

	assert(t, d.Get(`a.b.c.d.e`), nil)

	d.Set(`a.b.c.d.e`, `ooo`)
	assert(t, d.Get(`a.b.c.d.e`), `ooo`)

	d.Set(`a.b.c.d.e`, 4324)
	assert(t, d.Get(`a.b.c.d.e`), 4324)

	d.Set(78, 88)
	assert(t, d.Get(78), 88)

	d.Forget(78)
	assert(t, d.Get(78), nil)
}

//	test parse from json
func TestDict_ParseJson(t *testing.T) {
	d := NewDict()
	assert(t, len(d.data), 0)
	d.ParseJsonObject([]byte(`{"name":"zhr","age":18,"address":["yuncheng","beijing"]}`))
	assert(t, len(d.data), 3)

	err := d.ParseJsonObject([]byte(`["a","b"]`))
	assert(t, len(d.data), 0)
	assert(t, err, ErrNotDict)

}

//	test .Json() .Keys() .Values()
func TestDict_Json(t *testing.T) {
	d := NewDict()

	d.Set(1, 3)
	d.Set(`a`, 'b')
	d.Set(`a.b.c.d`, 'e')

	assert(t, d.Json(), `{"1":3,"a":{"b":{"c":{"d":101}}}}`)
	assert(t, d.Keys(), []string{`1`, `a`})
	assert(t, d.Values(), []interface{}{
		3,
		map[string]interface{}{
			`b`: map[string]interface{}{
				`c`: map[string]interface{}{
					`d`: 'e',
				},
			},
		},
	})
}

//	assert a equals b, or show code where error
func assert(t *testing.T, a interface{}, b interface{}) {
	if fmt.Sprint(a) != fmt.Sprint(b) {
		file, line := calledBy()
		t.Errorf(
			"Failure in %s:%d\nexpect:\t%v\n   but:\t%v\n----\n%s\n",
			file, line,
			a, b,
			showFile(file, line),
		)
	}
}

func calledBy() (string, int) {
	_, file, line, _ := runtime.Caller(2)
	return file, line
	file = strings.TrimPrefix(file, os.Getenv(`GOPATH`))
	return `$GOPATH` + file, line
}

func showFile(filename string, line int) string {
	file, _ := ioutil.ReadFile(filename)
	buf := bytes.NewBufferString(``)

	size := len(file)
	begin := line - 2
	end := line + 3

	//	@todo line超过文件时
	found := 1
	for idx := 0; idx < size; idx++ {
		if found >= begin && found < end {
			buf.WriteByte(file[idx])
		}

		if file[idx] == '\n' {
			found++

			if found >= begin && found < end {
				tag := ' '
				if found == line {
					tag = '#'
				}
				buf.WriteString(fmt.Sprintf("%c%4d|\t", tag, found))
			}

			if found >= end {
				break
			}
		}
	}

	return buf.String()
}
