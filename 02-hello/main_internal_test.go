package main

import "testing"

func ExampleMain() {
	main()
}

func TestGreet(t *testing.T) {
	type testCase struct {
		lang language
		want string
	}

	var tests = map[string]testCase{
		"English":    {lang: "en", want: "Hello world"},
		"French":     {lang: "fr", want: "Bonjour le monde"},
		"Akkadian":   {lang: "akk", want: `Unsupported language: "akk"`},
		"Greek":      {lang: "el", want: "Χαίρετε Κόσμε"},
		"Hebrew":     {lang: "he", want: "םלוע םולש"},
		"Urdu":       {lang: "ur", want: "ایند ولیہ"},
		"Vietnamese": {lang: "vi", want: "Xin chào Thế Giới"},
		"Empty":      {lang: "", want: `Unsupported language: ""`},
	}

	// range over all the scenarios
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := greet(tc.lang)

			if got != tc.want {
				t.Errorf("Expected: %q, got: %q", tc.want, got)
			}
		})
	}
}
