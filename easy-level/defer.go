package main

import "fmt"

// Что выведет код и почему?
func main() {
	fmt.Println("start")
	for i := 1; i < 4; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("end")
}

Ответ:
Выполнится вывод такого вида:
start
end
3
2
1
Так как отложенная функция defer сохраняется в стек функций, который работает по принципу LIFO(Last In First Out),
а значит первыми выйдут данные, что вошли последними
