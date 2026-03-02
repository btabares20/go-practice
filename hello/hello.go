package hello 

var translations = map[string]string{
	"english":"Hello",
	"spanish":"Hola",
	"french": "Bonjour",
}
// A function that takes and greets a name
func Hello(name, language string)string{
	if name == "" {
		name = "World"
	}
	return translations[language] + ", " + name
}
