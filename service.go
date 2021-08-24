package main

// Интерфейс для реализации многоуровневых сервисов
type MyService interface {
	Foo(info ResponseInfo, c chan int)
}
