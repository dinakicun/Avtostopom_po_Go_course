ЗАДАЧА:
package main
import (
	"fmt"
	"sync"
)
// time.Sleep использовать нельзя. это будет не валидным ответом на собеседовании
// Расскажи подробно что происходит
// Как сделать так, чтобы работало оставив канал небуферезированным?
func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(v int) {
			defer wg.Done()
			ch <- v * v
		}(i)
	}
	wg.Wait()
	var sum int
	for v := range ch {
		sum += v
	}
	fmt.Printf("result: %d\n", sum)
}

РЕШЕНИЕ: 
Три горутины отправляют в небуферизованный канал данные. 
main вызывает wg.Wait() и только потом читает канал через range и суммирует значения.
Но программа упадет с deadlock, потому что каналам необходим получатель, но из-за wg.Wait() они застряли
и не дойдут до читателей в main

Вторая проблема в том, что канал нигде не закрывается, а range завершается только по закрытию канала, т.е. он тоже зависнет

Для этого нам нужно вынести wg.Wait() и закрытие канала в другую горутину, чтобы они шли одновременно.

func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(v int) {
			defer wg.Done()
			ch <- v * v
		}(i)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	var sum int
	for v := range ch {
		sum += v
	}
	fmt.Printf("result: %d\n", sum)
}
