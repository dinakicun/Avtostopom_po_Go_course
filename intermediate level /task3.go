package main

import (
	"fmt"
	"sync"
	"time"
)

// Расскажи подробно что происходит
Программа симулирует отправку 10000 запросов практически одновременно
и считает их количество, после чего выводит
Выведется ровно 10000, т.к. мы закрыли одновременный доступ к count через Mutex и мы ждем
конца выполнения всех горутин благодаря WaitaGroup
// Сколько времени будет выполняться?
Пока одна горутина спит, планировщик го запускает следующую горутину 
Поэтому программа выполнится очень быстро за чуть больше, чем милисекунду

const numRequests = 10000
var count int
var m sync.Mutex

func networkRequest() {
	time.Sleep(time.Millisecond) // Эмуляция сетевого запроса.
	m.Lock()
	count++
	m.Unlock()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(numRequests)
	for i := 0; i < numRequests; i++ {
		go func() {
			defer wg.Done()
			networkRequest()
		}()
	}

	wg.Wait()
	fmt.Println(count)
}
