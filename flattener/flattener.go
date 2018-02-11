package flattener

type Flattener interface {
	Flatten(data []byte) (map[string]interface{}, error)
}
