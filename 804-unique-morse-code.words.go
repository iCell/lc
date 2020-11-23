var morses = []string{
	".-", 
	"-...", 
	"-.-.", 
	"-..", 
	".", 
	"..-.", 
	"--.", 
	"....", 
	"..", 
	".---", 
	"-.-", 
	".-..", 
	"--", 
	"-.", 
	"---", 
	".--.", 
	"--.-", 
	".-.", 
	"...", 
	"-", 
	"..-", 
	"...-", 
	".--", 
	"-..-", 
	"-.--", 
	"--..",
}

func uniqueMorseRepresentations(words []string) int {
	unique := make(map[string]bool)
	for _, word := range words {
		var encoded strings.Builder
		for _, c := range word {
			encoded.WriteString(morses[c-'a'])
		}
		unique[encoded.String()] = true
	}
	return len(unique)
}