package Collection

type self Collection

type mapFunc func(interface{}) interface{}

type Collection interface {
	Get() []interface{}

	Push() self

	Map(mapFunc) self

	ToJson() []byte
}
