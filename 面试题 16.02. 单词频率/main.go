package main

type WordsFrequency struct {
	books map[string]int
}

func Constructor(book []string) WordsFrequency {
	m := map[string]int{}

	for _, v := range book {
		m[v]++
	}
	return WordsFrequency{m}
}

func (this *WordsFrequency) Get(word string) int {

	return this.books[word]
}

func main() {

}
