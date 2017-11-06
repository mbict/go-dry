package caseing

// SnakeCase produces the snake_case of a CamelCased string
func SnakeCase(name string) string {
	return xCase(name, '_')
}
