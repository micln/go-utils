package dict

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/micln/go-utils"
)

var (
	ErrNotDict = errors.New(`parsed object is not map[string]interface{}`)
)

type dict struct {
	provider map[string]interface{}
}

func Dict() *dict {
	return &dict{
		provider: newMap(),
	}
}

func (d *dict) Get(k interface{}) interface{} {
	paths := strings.Split(toString(k), `.`)

	var current interface{}
	current = toMap(d.provider)

	size := len(paths)
	for i := 0; i < size-1; i++ {
		m := toMap(current)
		current = m[paths[i]]
	}

	return toMap(current)[paths[size-1]]
}

func (d *dict) Set(k interface{}, v interface{}) {
	paths := strings.Split(toString(k), `.`)

	parent := d.provider

	size := len(paths)
	for idx := 0; idx < size-1; idx++ {
		//fmt.Println(idx, d.provider, parent, &d.provider, &parent)
		seq := paths[idx]

		i := parent[seq]
		if _, ok := i.(map[string]interface{}); !ok {
			parent[seq] = newMap()
			parent = parent[seq].(map[string]interface{})
		} else {
			parent = i.(map[string]interface{})
		}
	}

	parent[paths[size-1]] = v
}

func (d *dict) Forget(k interface{}) {
	d.Set(k, nil)
}

func (d *dict) ParseJson(jsonBytes []byte) (err error) {
	var i interface{}
	err = json.Unmarshal(jsonBytes, &i)
	if err != nil {
		return
	}

	var ok bool
	if d.provider, ok = i.(map[string]interface{}); !ok {
		return ErrNotDict
	}

	return
}

func (d *dict) Keys() (keys []string) {
	for k := range d.provider {
		keys = append(keys, k)
	}
	return
}

func (d *dict) Values() (values []interface{}) {
	for _, v := range d.provider {
		values = append(values, v)
	}
	return
}

func (d *dict) Json() string {
	return go_utils.JsonEncode(d.provider)
}

func toString(k interface{}) string {
	switch k.(type) {
	case string:
		return k.(string)
	case int:
		return strconv.FormatInt(int64(k.(int)), 10)
	default:
		return fmt.Sprint(k)
	}
}

func newMap() map[string]interface{} {
	return make(map[string]interface{})
}

func toMap(i interface{}) map[string]interface{} {
	m, ok := i.(map[string]interface{})
	if !ok {
		//fmt.Printf("%v,%v\n", i, m)
		m = make(map[string]interface{})
		//fmt.Printf("%v,%v\n", i, m)
	}

	return m
}
