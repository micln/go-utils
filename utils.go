package go_utils

/**
常用小函数
Go 标准库的接口太底层了, 不适用于应用开发。这里稍作封装
*/

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
	"bytes"
	"strconv"
)

func JsonEncode(s interface{}) string {
	b, _ := json.Marshal(s)
	return string(b)
}

func JsonDecode(s interface{}) (v interface{}) {
	if ss, ok := s.(string); ok {
		s = []byte(ss)
	}
	json.Unmarshal(s.([]byte), &v)
	return
}

func Base64Decode(s string) string {
	b, _ := base64.StdEncoding.DecodeString(s)
	return string(b)
}

func Base64EncodeString(s string) string {
	return Base64Encode([]byte(s))
}

func Base64Encode(bs []byte) string {
	return base64.StdEncoding.EncodeToString(bs)
}

//  @return [0, max)
func Rand(max int64) int64 {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(max)
}

func Date(format string, timestamp ...int64) string {
	seconds := time.Now().Local().Unix()
	if len(timestamp) > 0 {
		seconds = timestamp[0]
	}

	rule := []string{
		`Y`, `2006`, //  年
		`y`, `06`,

		`m`, `01`, //  月
		`n`, `1`,

		`d`, `02`, //  日
		`j`, `2`,

		`H`, `15`, //  时
		`h`, `03`,
		`g`, `3`,
		`G`, `15`, // 应该木有前导零

		`i`, `04`, //  分
		`s`, `05`, //  秒

		`D`, `Mon`, //  周
		`N`, `1`,
	}

	specs := func(expr string) string {
		switch expr {
		case `S`:
			return `st/nd/rd/th`
		case `z`:
			return `The day of the year`
		case `t`:
			return `Number of days in the given month`
		case `L`:
			return `Whether it's a leap year`
		case `a`:
			return `am or pm`
		case `A`:
			return `AM or PM`
		default:
			return expr
		}
	}

	size := len(rule)
	for i := 0; i < size; i += 2 {
		format = strings.Replace(format, rule[i], rule[i+1], -1)
	}

	format = specs(format)

	return time.Unix(seconds, 0).Local().Format(format)
}

func StrToTime(expr string) int64 {
	t := time.Now().Local()
	//pattern := `\d+ [(year)|(month)|(day)|(hour)|(minute)|(second)]`

	for i := 1; i < 10; i++ {

	}
	return t.Unix()
}

func Md5(s string) string {
	return fmt.Sprintf(`%x`, md5.Sum([]byte(s)))
}

func Sha1(s string) string {
	return fmt.Sprintf(`%x`, sha1.Sum([]byte(s)))
}

type ErrBar struct{}

var Err ErrBar

func (e *ErrBar) Panic(err error) {
	if err != nil {
		panic(err)
	}
}

func (e *ErrBar) Fatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func (e *ErrBar) Log(err error) {
	if err != nil {
		log.Println(err)
	}
}

// 兼容 func or value
//func value(v interface{}) interface{} {
//	if f, ok := v.(func() {}); ok {
//		return f();
//	}
//
//	return v;
//}


func InSlice(v interface{}, sl ...interface{}) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

//	convert `a_b_c` to `aBC`
func CamelCase(v string) string{
	ret := make([]byte, 31)
buf := bytes.NewBuffer(ret)
	length := len(v)
	for i:=0;i<length;i++{
		if v[i]!='_' {
			buf.WriteByte(v[i])
		}else{
			i+=1
			if i>length{
				continue
			}
			if v[i] >= 'a' && v[i]<='z' {
				buf.WriteByte(v[i]+ 'A'-'a')
			}else{
			buf.WriteByte(v[i])
			}
		}
	}

	return buf.String()
}

//	convert `ABC` to `a_b_c`
func UnderlineCase(v string) string{
	ret := make([]byte, 31)
	buf := bytes.NewBuffer(ret)
	length := len(v)
	for i:=0;i<length;i++{
		if v[i] > 'A' && v[i]<'Z'{
			if i!=0{
				buf.WriteByte('_')
				buf.WriteByte(v[i] + 'a' -'A')
			}
		}
	}
	return buf.String()
}

func Intval(s string) int {
	i,_:=strconv.Atoi(s)
	return i
}