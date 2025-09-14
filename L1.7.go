package main

import "sync"

// структура, которая хранит map и мьютекс
type SafeMap struct {
	mu sync.Mutex     // мьютекс, который защищает доступ к карте
	m  map[string]int // сама карта (ключ - строка, значение - число)
}

// функция-конструктор: создаём новую структуру SafeMap
func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[string]int), // инициализируем пустую map
	}
}

// метод записи: кладём значение по ключу
func (s *SafeMap) Set(key string, value int) {
	s.mu.Lock()      // блокируем мьютекс (доступ только одной горутине)
	s.m[key] = value // пишем в map
	s.mu.Unlock()    // разблокируем мьютекс
}

// метод чтения: достаём значение по ключу
func (s *SafeMap) Get(key string) (int, bool) {
	s.mu.Lock()       // блокируем мьютекс
	v, ok := s.m[key] // читаем значение и признак наличия ключа
	s.mu.Unlock()     // разблокируем мьютекс
	return v, ok      // возвращаем результат
}
