package main

import (
	"context"
	"fmt"
	"time"
)

// функция отправитель
func Sender(ctx context.Context, ch chan<- int) {
	i := 0 // инициализируем начальное значение переменной
	for {
		select {
		case <-ctx.Done(): // кейс, когда время вышло
			return
		case ch <- i: // кейс, когда отправляем значение в канал
			i++                                // увеличиваем значение на 1
			time.Sleep(time.Millisecond * 200) // пауза для вывода
		}
	}
}

// функция получатель
func Receiver(ctx context.Context, ch <-chan int) {
	for {
		select {
		case <-ctx.Done(): // кейс, когда время вышло
			return
		case value, ok := <-ch: // кейс, когда получаем из канала значение при условии, что канал не закрыт
			if !ok {
				return
			}
			fmt.Println(value) // печатаем значение
		}
	}
}
