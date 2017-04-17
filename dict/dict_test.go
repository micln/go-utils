package dict

import (
	"strconv"
	"testing"

	. "github.com/micln/go-utils/test"
)

//	test Getter/Setter
func TestNewDict(t *testing.T) {
	d := NewDict()
	AssertEqual(t, d.Get(`a.b.c.d.e`), nil)

	d.Set(`a.b.c.d.e`, `ooo`)
	AssertEqual(t, d.Get(`a.b.c.d.e`), `ooo`)

	d.Set(`a.b.c.d.e`, 4324)
	AssertEqual(t, d.Get(`a.b.c.d.e`), 4324)

	d.Set(78, 88)
	AssertEqual(t, d.Get(78), 88)

	d.Forget(78)
	AssertEqual(t, d.Get(78), nil)
}

func TestDict_Filter(t *testing.T) {
	d1 := NewDict()
	for i := 1; i < 10; i++ {
		d1.Set(i, i*i)
	}

	AssertEqual(t, d1.Len(), 9)

	d2 := d1.Filter(func(v interface{}, k string) bool {
		i, _ := strconv.Atoi(k)
		return i > 5
	})

	AssertEqual(t, d2.Len(), 4)

	d3 := d1.Filter(func(v interface{}, k string) bool {
		return v.(int) < 10
	})

	AssertEqual(t, d3.Len(), 3)
}

//	test parse from json
func TestDict_ParseJson(t *testing.T) {
	d := NewDict()
	AssertEqual(t, len(d.data), 0)
	d.ParseJsonObject([]byte(`{"name":"zhr","age":18,"address":["yuncheng","beijing"]}`))
	AssertEqual(t, len(d.data), 3)

	err := d.ParseJsonObject([]byte(`["a","b"]`))
	AssertEqual(t, len(d.data), 0)
	AssertEqual(t, err, ErrNotDict)

}

//	test .Json() .Keys() .Values()
func TestDict_Json(t *testing.T) {
	d := NewDict()

	d.Set(1, 3)
	d.Set(`a`, 'b')
	d.Set(`a.b.c.d`, 'e')

	AssertEqual(t, d.Json(), `{"1":3,"a":{"b":{"c":{"d":101}}}}`)
	AssertEqual(t, d.Keys(), []string{`1`, `a`})
	AssertEqual(t, d.Values(), []interface{}{
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
