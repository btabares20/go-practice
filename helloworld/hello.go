package helloworld


func Hello(name, lang string) string {
	langs := map[string]string{
		"english": "Hello, ",
		"spanish": "Hola, ",
	}
	if lang == "" {
		lang = "english"
	}
	if name == "" {
		name = "World"
	}
	if _, exists := langs[lang]; !exists {
		lang = "english"
	}
	hello := langs[lang]
	return hello + name
}
