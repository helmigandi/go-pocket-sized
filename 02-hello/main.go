package main

import (
	"flag"
	"fmt"
)

// Read from input cmd: go run main.go -lang=el
// To test in cmd: go test
func main() {
	var lang string
	flag.StringVar(&lang, "lang", "en", "The required language e.g. en, fr, ur, ...")
	flag.Parse()

	greeting := greet(language(lang))

	fmt.Println(greeting)
}

// language represents the language’s code
type language string

// phrasebook holds greeting for each supported language
var phrasebook = map[language]string{
	"el": "Χαίρετε Κόσμε",     // Greek
	"en": "Hello world",       // English
	"fr": "Bonjour le monde",  // French
	"he": "םלוע םולש",         // Hebrew
	"ur": "ایند ولیہ",         // Urdu
	"vi": "Xin chào Thế Giới", // Vietnamese
}

// greet says hello to the world in the specified language
func greet(l language) string {
	greeting, ok := phrasebook[l]

	if !ok {
		return fmt.Sprintf("Unsupported language: %q", l)
	}

	return greeting
}
