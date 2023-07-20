package hello

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Ola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {

	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "spanish":
		prefix = spanishHelloPrefix

	case "french":
		prefix = frenchHelloPrefix

	default:
		prefix = englishHelloPrefix
	}
	return prefix

}

func main() {
	// fmt.Println(Hello("world"))
}
