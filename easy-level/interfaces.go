package main

import "fmt"

// Что выведет код и почему?
type impl struct{}

type I interface {
	C()
}

func (*impl) C() {}

func A() I {
	return nil

}
func B() I {
	var ret *impl
	return ret
}

func main() {
	a := A()
	b := B()
	fmt.Println(a == b)
}

ОТВЕТ: код выведет false, т.к в функции A вернется интерфейс, равный nil(и тип, и значения не определены)
а функция B вернет интерфейс, у которого тип равен *impl, а значение равно nil
