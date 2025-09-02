package main

import (
	"goroutineGolang-wb/model"
)

func main() {
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
