package Collection

import "container/list"

type self *ArrayCollector
type ArrayCollector struct {
	items []interface{}
}

func (arr self) Get() self {
	list.New()
}
