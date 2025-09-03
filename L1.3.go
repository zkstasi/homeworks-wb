package main

import (
	"context"
	"fmt"
)

func Channels(N int, ctx context.Context) {

	ch := make(chan int) // создаем канал

	// запускаем N воркеров
	for i := 0; i < N; i++ {
		go worker(ch, i+1, ctx) // запускаем воркера как горутину с идентификатором воркера i + 1
	}

	// пишем данные в канал постоянно в бесконечном цикле
	i := 0
	for {
		select {
		case <-ctx.Done(): // если пришел сигнал отмены, завершаем запись
			close(ch) // закрываем канал, чтобы воркеры, читающие через range, завершились
			return
		default:
			ch <- i
			i++
		}
	}
}

func worker(ch chan int, id int, ctx context.Context) {
	for {
		select {
		case value, ok := <-ch: // читаем данные из канала
			if !ok {
				// Канал закрыт, завершаем работу
				return
			}
			// печатаем данные из канала и какой воркер прочитал данные
			fmt.Printf("worker %d: %d\n", id, value)
		case <-ctx.Done():
			// Получен сигнал отмены, завершаем работу
			return
		}
	}
}
