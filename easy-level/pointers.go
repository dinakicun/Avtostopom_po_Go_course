package main

import "fmt"

// Что выведет код и почему?
// Как изменить v на 3 через функцию без добавления return из функции
func main() {
	v := 5
	p := &v
	fmt.Println(*p)

	changePointer(p)
	fmt.Println(*p)
}

func changePointer(p *int) {
	v := 3
	p = &v
}

Код выведет:
5
5

Так как мы меняем содержимое переменной p, существующей только в функции
Это просто копия нашего указателя
Для того, чтобы изменить именно значение p в main,
необходимо разыменовать указатель и изменить значение по адресу

func main() {
	v := 5
	p := &v
	fmt.Println(*p)

	changePointer(p)
	fmt.Println(*p)
}

func changePointer(p *int) {
	v := 3
	*p = v
}
