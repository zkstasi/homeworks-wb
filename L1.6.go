package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// Выход по условию
func stopByCondition() {

	// запускаем горутину
	go func() {

		// цикл ограничен условием
		for i := 0; i < 10; i++ {
			fmt.Println("Горутина работает, i =", i)
			time.Sleep(200 * time.Millisecond)
		}

		// цикл закончился и горутина завершается
		fmt.Println("Горутина завершена по условию")
	}()

	// ждем, чтобы горутина отработала
	time.Sleep(3 * time.Second)
}

// Остановка через канал
func stopByChannel() {

	// создаем канал для сигналов
	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Горутина завершена через done-cnannel")
				return
			default:
				fmt.Println("Горутина работает...")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	// ждем
	time.Sleep(1 * time.Second)
	// шлем сигнал остановки в канал
	done <- struct{}{}

	// ждем, чтобы увидеть сообщение о завершении
	time.Sleep(300 * time.Millisecond)
}

// Остановка через контекст
func stopByContext() {

	// создаем контекст
	ctx, cancel := context.WithCancel(context.Background())
	// откладываем вызов cancel()
	defer cancel()

	// в горутине бесконечный цикл проверяет наличие сигнала отмены
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done(): // кейс получения сигнала для завершения работы горутины
				fmt.Println("Горутина завершена через контекст")
				return
			default: // если сигнала нет, то работа продолжается
				fmt.Println("Горутина работает...")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}(ctx)

	// ждем, чтобы корректно сработал defer, а он сработает после завершения функции
	time.Sleep(1 * time.Second)
}

// Остановка через Runtime.Goexit()
func stopByGoexit() {

	// запускаем цикл в горутине
	go func() {
		for i := 0; i < 7; i++ {
			fmt.Println("Работаю", i) // печатаем текущее значение счетчика

			// условие, при котором остановится горутина
			if i == 3 {
				fmt.Println("Вызов Goexit()")
				runtime.Goexit() // немедленно завершает горутину
			}
			time.Sleep(200 * time.Millisecond) // пауза для корректного вывода
		}
	}()
	time.Sleep(2 * time.Second) // ждем, чтобы горутина показала работу
}

// Остановка через закрытия канала
func stopByCloseChannel() {
	ch := make(chan int) // канал для передачи данных

	go func() { // запускаем горутину, получающую данные
		for value := range ch { // читаем из канала, пока он открыт
			fmt.Println("Получено значение:", value)
		}
		fmt.Println("Горутина завершена после закрытия канала") // выходим, канал закрыт
	}()

	// отправляем данные в канал в цикле
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(200 * time.Millisecond) // пауза для корректного вывода
	}

	close(ch)                          // закрытие канала
	time.Sleep(400 * time.Millisecond) // ждем отработки горутины
}
