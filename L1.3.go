package main

import (
	"fmt"
)

func Channels(N int) {

	ch := make(chan int) // создаем канал

	// запускаем N воркеров
	for i := 0; i < N; i++ {
		go worker(ch, i+1) // запускаем воркера как горутину с идентификатором воркера i + 1
	}

	// пишем данные в канал постоянно в бесконечном цикле
	i := 0
	for {
		ch <- i
		i++
	}
}

func worker(ch chan int, id int) {

	// воркер читает данные из канала
	for value := range ch {

		// печатаем данные из канала и какой воркер прочитал данные
		fmt.Printf("worker %d: %d\n", id, value)
	}
}
