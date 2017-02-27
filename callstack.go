package go_utils

import (
	"bytes"
	"fmt"
	"runtime"
)

type CallStackNode struct {
	Pc   uintptr
	File string
	Line int
}

func (c *CallStackNode) String() string {
	return fmt.Sprintf("%s:%d", c.File, c.Line)
}

type CallStacks []*CallStackNode

func (cs CallStacks) String() string {
	buf := bytes.NewBufferString(``)
	for i := range cs {
		buf.WriteString(cs[i].String())
		buf.WriteByte('\n')
	}
	return buf.String()
}

func GetCallStacks() CallStacks {
	stacks := []*CallStackNode{}
	for i := 0; i < 50; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}

		stacks = append(stacks, &CallStackNode{
			Pc:   pc,
			File: file,
			Line: line,
		})
	}
	return stacks
}
