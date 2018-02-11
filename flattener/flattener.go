package flattener

// Flattener is the interface that wraps the data flattening method.
type Flattener interface {
	// Flatten transforms nested data structure into flat key, value pars
	Flatten(data []byte) (map[string]interface{}, error)
}
