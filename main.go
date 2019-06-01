package main

func main() {
	text := "ababa"
	words := []string{"aba", "ab"}
	res := indexPairs(text, words)
	for _, pair := range res {
		println(pair[0], pair[1])
	}
}
