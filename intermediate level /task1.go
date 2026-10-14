ЗАДАЧА: 
package main
import "fmt"

// time.Sleep использовать нельзя. это будет не валидным ответом на собеседовании
// 1. Как будет работать код?
// 2. Как сделать так, чтобы выводился только первый ch?
// 2.1 Удалять из существующего кода ничего нельзя. Только добавлять)
func main() {
	ch := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)
	go func() {
		ch <- true
	}()
	go func() {
		ch2 <- true
	}()
	go func() {
		ch3 <- true
	}()

	select {
	case <-ch:
		fmt.Printf("val from ch")
	case <-ch2:
		fmt.Printf("val from ch2")
	case <-ch3:
		fmt.Printf("val from ch3")
	}
}


РЕШЕНИЕ:
1. Выведется одна строка, т.к. select выполняет только одну ветку и завершается. Горутины запускаются практически одновременно
и не предугадать какая выполнится быстрее, а возможно они уже все будут готовы до того, как
программа дойдент до select и выберется случайная
2. Чтобы вывелся только первый ch, я перенесла вызов других горутин в первую,чтобы горутины для ch2 и ch3 
стартовали только после того, как select уже выбрал ch.
package main
import "fmt"

func main() {
	ch := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)
	go func() {
		ch <- true
    		go func() {
    			ch2 <- true
    		}()
    		go func() {
    			ch3 <- true
    		}()
	}()

	select {
	case <-ch:
		fmt.Printf("val from ch")
	case <-ch2:
		fmt.Printf("val from ch2")
	case <-ch3:
		fmt.Printf("val from ch3")
	}
}
