package main

import (
	"fmt"
)

// 1
// Что происходит?
// Как сделать так, чтобы работало?
func main() {
	var m map[string]int
	for _, word := range []string{"hello", "world", "from", "the",
		"best", "language", "in", "the", "world"} {
		m[word]++
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

Ответ: программа пытается высчитать количество повторений слов в слайсе, 
используя мапы.
Но программа упадет с паникой, так как мапа m равна nil (не была инициализирована через make)

Исправленный вариант:
func main() {
	m := make(map[string]int, 9)
	for _, word := range []string{"hello", "world", "from", "the",
		"best", "language", "in", "the", "world"} {
		m[word]++
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
