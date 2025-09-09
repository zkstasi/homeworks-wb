package main

import (
	"fmt"
	"time"
)

func main() {
	// Person() // вызов функции для L1.1
	// GoroutineVersionOne() // вызов функции для L1.2
	// GoroutineVersionTwo() // вызов функции для L1.2

	// запустить программу нужно из терминала с помощью команды go run main.go L1.3.go 5, где 5 - кол-во воркеров
	/*
		// Проверяем, передан ли аргумент командной строки (количество воркеров)
		if len(os.Args) < 2 {
			// Если аргумент не передан, выводим инструкцию по запуску программы
			fmt.Println("Usage: go run main.go L1.3.go <num_workers>")
			return // Завершаем выполнение main(), так как нет нужного аргумента
		}

		// Преобразуем первый аргумент командной строки в целое число
		N, err := strconv.Atoi(os.Args[1])

		// Проверяем, была ли ошибка преобразования или число меньше, или равно 0
		if err != nil || N <= 0 {

			// Если аргумент некорректный, выводим сообщение об ошибке
			fmt.Println("Введите корректное число воркеров")
			return // Завершаем выполнение main(), так как введено неправильное значение
		}

		//создание контекста, который отменится, когда пользователь нажмет Ctrl+C или придет другой сигнал завершения
		ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer stop() // освобождаем ресурсы

		Channels(N, ctx) // вызываем функцию из L1.3.go
	*/

	/*
		// время работы программы
		N := 3

		// контекст с таймаутом N секунд
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(N)*time.Second)
		defer cancel()

		// создаём канал для передачи данных
		ch := make(chan int)

		// запускаем горутину-отправителя
		go Sender(ctx, ch)

		// запускаем горутину-получателя
		go Receiver(ctx, ch)

		// ждём завершения контекста
		<-ctx.Done()

		// закрываем канал (сигнализируем получателю, что данных больше не будет)
		close(ch)
	*/

	fmt.Println("Остановка по условию")
	stopByCondition()       // вызов функции
	time.Sleep(time.Second) // ждем, чтобы горутина успела завершится

	fmt.Println("Остановка через канал уведомления")
	stopByChannel()         // вызов функции
	time.Sleep(time.Second) // ждем, чтобы горутина успела завершится

	fmt.Println("Остановка через контекст")
	stopByContext()         // вызов функции
	time.Sleep(time.Second) // ждем, чтобы горутина успела завершится

	fmt.Println("Остановка через runtime.Goexit()")
	stopByGoexit()          // вызов функции
	time.Sleep(time.Second) // ждем, чтобы горутина успела завершится

	fmt.Println("Остановка через закрытие канала")
	stopByCloseChannel()    // вызов функции
	time.Sleep(time.Second) // ждем, чтобы горутина успела завершится

	fmt.Println("\nВсе демонстрации завершены.")
}
