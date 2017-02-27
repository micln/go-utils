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

func TestNewDict(t *testing.T) {
	d := Dict()

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

func TestDict_ParseJson(t *testing.T) {
	d := Dict()
	assert(t, len(d.provider), 0)
	d.ParseJson([]byte(`{"name":"zhr","age":18,"address":["yuncheng","beijing"]}`))
	assert(t, len(d.provider), 3)

	err := d.ParseJson([]byte(`["a","b"]`))
	assert(t, len(d.provider), 0)
	assert(t, err, ErrNotDict)

}

func TestDict_Json(t *testing.T) {
	d := Dict()

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
