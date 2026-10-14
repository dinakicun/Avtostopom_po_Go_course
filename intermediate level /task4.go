package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.NewSource(time.Now().UnixNano()) //для составления рандомных чисел. больше тебе знать не нужно
}

// time.Sleep использовать нельзя. это будет не валидным ответом на собеседовании
// Есть функция unpredictableFunc, работающая неопределенно долго и возвращающая число.
// Её тело нельзя изменять (представим, что внутри сетевой запрос).
func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000) //рандомное число
	time.Sleep(time.Duration(rnd) * time.Millisecond)

	return rnd
}

// Нужно изменить функцию обертку predictableFunc, которая будет работать с заданным таймаутом (например, 1 секунду).
// Если "длинная" функция отработала за это время - отлично, возвращаем результат.
// Если нет - возвращаем ошибку. Результат работы в этом случае нам не важен.
РЕШЕНИЕ: 
func predictableFunc() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	ch := make(chan int64)

	go func() {
		ch <- unpredictableFunc()
	}()

	select {
	case <-ctx.Done():
		return 0, fmt.Errorf("waited for too long")
	case result := <-ch:
		return result, nil
	}
}


func main() {
	fmt.Println("started")

	fmt.Println(predictableFunc())
}
