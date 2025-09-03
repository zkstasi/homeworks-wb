package main

import (
	"fmt"
	"goroutineGolang-wb/model"
	"sync"
)

func Human() {
	// Создаём объект Action с ID=1 и встроенным Human
	a := model.Action{
		ID: 1,
		Human: model.Human{
			Name:       "Анастасия",
			NameDative: "Анастасии",
			Age:        29,
			Gender:     "Женщина",
			Phone:      "555-555-5555",
			Email:      "anastasia@gmail.com",
			Address:    "улица Ленина дом 1",
		},
	}

	a.SayHello()               // вызываем метод Human через embedding, здороваемся
	a.ContactInfo()            // выводим контактные данные
	a.Eat("апельсин")          // человек ест мясо
	a.Birthday()               // увеличиваем возраст на 1
	a.Sleep(8)                 // показываем, что человек спит 8 часов
	a.StartActivity("jogging") // начинаем бегать трусцой
	a.ShowActivity()           // Демонстрируем текущую активность при наличии
	a.StopActivity()           // останавливаемся
}

// первый вариант решения задачи
func GoroutineVersionOne() {
	// Human()

	// объявляем массив
	arr := [5]int{2, 4, 6, 8, 10}

	// создаем WaitGroup для корректной работы горутин и основной горутины main
	var wg sync.WaitGroup

	// в цикле запускаем функцию-горутину, где каждая итерация запускает отдельную горутину для одного элемента
	for i := 0; i < len(arr); i++ {

		// увеличиваем счётчик ожидаемых горутин на 1
		wg.Add(1)

		// вызов функции Kvadrat через оператор go
		go Kvadrat(arr[i], &wg)
	}
	// ждем завершения горутин
	wg.Wait()
}

// функция возведения числа в квадрат, принимает число и указатель на WaitGroup
func Kvadrat(i int, wg *sync.WaitGroup) {

	// уменьшает счётчик WaitGroup при завершении горутины
	defer wg.Done()

	// печатаем квадрат числа
	fmt.Println(i * i)
}

// оптимизированный вариант решения задачи
func GoroutineVersionTwo() {

	// объявляем и инициализируем массив
	arr := [5]int{2, 4, 6, 8, 10}

	// объявляем WaitGroup для того, чтобы программа дождалась завершения горутин
	var wg sync.WaitGroup

	// перебираем в цикле все элементы массива, игнорируя индексы
	for _, num := range arr {

		// увеличиваем счётчик ожидаемых горутин на 1
		wg.Add(1)

		// запуск горутины
		// анонимная функция, которая принимает число m
		go func(m int) {

			// defer гарантирует, что после завершения горутины WaitGroup уменьшит счетчик
			defer wg.Done()

			// выводим квадрат числа
			fmt.Println(m * m)
		}(num) // передаем текущее число из цикла в параметр m, чтобы избежать замыкания
	}

	// ждем пока все горутины завершатся
	wg.Wait()
}

func main() {
	// Human()
	//GoroutineVersionOne()
	GoroutineVersionTwo() // вызов функции
}
