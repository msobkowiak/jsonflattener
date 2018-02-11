package flattener

type Flattener interface {
	Flatten(data []byte, delimiter string) ([]byte, error)
}
