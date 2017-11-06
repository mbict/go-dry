package caseing

// KebabCase produces the kebab-case of a CamelCased string
func KebabCase(name string) string {
	return xCase(name, '-')
}
