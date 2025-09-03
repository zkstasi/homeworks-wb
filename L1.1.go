package main

import "fmt"

type Human struct {
	Name       string // Имя человека
	NameDative string // Имя человека в дательном падеже
	Age        int    // Возраст
	Gender     string // Пол
	Phone      string // Номер телефона
	Email      string // Адрес эл почты
	Address    string // "Земной" адрес
}

//  Методы структуры Human

func (h Human) SayHello() {
	fmt.Println("Привет, меня зовут", h.Name) // Человек здоровается
}

func (h Human) Walk() {
	fmt.Println(h.Name, "гуляет") // Человек гуляет
}

func (h Human) Learn() {
	fmt.Println(h.Name, "изучает Golang") // Человек изучает Golang
}

func (h Human) Eat(food string) {
	fmt.Println(h.Name, "ест", food) // Человек ест еду
}

func (h Human) Sleep(hours int) {
	fmt.Println(h.Name, "спит", hours, "часов") // Человек спит сколько-то часов
}

func (h *Human) Birthday() {
	h.Age++ // увеличиваем возраст на 1, чтобы возраст изменился используем указатель
	fmt.Println(h.NameDative, "исполнилось", h.Age)
}

func (h Human) ContactInfo() {
	// форматированный вывод информации о человеке с переводом на новую строку
	fmt.Printf("Name: %s\nPhone: %s\nAddress: %s\nEmail: %s\n", h.Name, h.Phone, h.Address, h.Email)
}

// Определение структуры Action
type Action struct {
	ID       int    // Идентификатор действия
	Activity string // Действие, чем занимается человек
	Human           // Встраиваем структуру Human (embedded struct)
}

// Методы структуры Action

func (a *Action) StartActivity(activity string) {
	a.Activity = activity // меняем текущую активность, поэтому используем указатель и выводим сообщение
	fmt.Println(a.Name, "начал(а) заниматься", activity)
}

// Показываем текущее действие (без изменения)
func (a Action) ShowActivity() {
	if a.Activity == "" {
		fmt.Println(a.Name, "ничем не занимается в данный момент")
	} else {
		fmt.Println(a.Name, "в настоящее время занимается:", a.Activity)
	}
}

func (a *Action) StopActivity() {
	fmt.Println(a.Name, "закончил(а) заниматься", a.Activity) // останавливаем активность и выводим сообщение
	a.Activity = ""
}

func Person() {
	// Создаём объект Action с ID=1 и встроенным Human
	a := Action{
		ID: 1,
		Human: Human{
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
