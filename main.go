package main

import (
	"fmt"
	"math/rand"
	"sync"
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

	/*
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
	*/

	rand.Seed(time.Now().UnixNano()) // инициализируем генератор случайных чисел текущим временем

	sm := NewSafeMap()    // создаём наш безопасный map
	var wg sync.WaitGroup // создаём группу ожидания для горутин

	// запускаем несколько горутин-писателей
	for i := 0; i < 3; i++ { // будет 3 горутины
		wg.Add(1)         // увеличиваем счётчик ожидаемых горутин
		go func(id int) { // анонимная функция, которая запускается как горутина
			defer wg.Done()           // по завершению уменьшаем счётчик
			for j := 0; j < 10; j++ { // пишем 10 значений
				key := fmt.Sprintf("key-%d-%d", id, j) // формируем ключ
				sm.Set(key, j)                         // кладём значение
				time.Sleep(time.Millisecond * 10)      // маленькая задержка, чтобы имитировать работу
			}
		}(i) // передаём номер горутины в функцию
	}

	// запускаем несколько горутин-читателей
	for i := 0; i < 3; i++ { // будет 3 горутины
		wg.Add(1) // увеличиваем счётчик
		go func(id int) {
			defer wg.Done()           // уменьшаем по завершению
			for j := 0; j < 10; j++ { // 10 попыток чтения
				key := fmt.Sprintf("key-%d-%d", rand.Intn(3), rand.Intn(10)) // случайный ключ
				if val, ok := sm.Get(key); ok {                              // пробуем прочитать
					fmt.Printf("Reader %d: key=%s val=%d\n", id, key, val) // если есть — выводим
				}
				time.Sleep(time.Millisecond * 10) // небольшая задержка
			}
		}(i)
	}

	wg.Wait()                       // ждём, пока все горутины закончат
	fmt.Println("Работа завершена") // выводим сообщение
}
