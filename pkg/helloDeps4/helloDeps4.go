package helloDeps4

import "fmt"

func PrintPhrase(phrase string) string {
	fmt.Println(phrase, "called helloDeps4.PrintPhrase")
	return phrase
}
