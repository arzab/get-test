package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

// Сервис для отправки http запросов
type clientService struct {
	http.Client
	MyService
}

func (service clientService) Foo(info ResponseInfo, c chan int) {
	// Засекаем время отправки запроса
	begin := time.Now()

	// Отправка запроса
	resp, err := service.Get(info.URL)

	info.Duration = time.Since(begin)

	if err != nil {
		info.Err = err.Error()
	} else {
		// Получаем код статуса
		info.StatusCode = resp.StatusCode

		// Получаем размер документа
		body, _ := ioutil.ReadAll(resp.Body)
		len := len(body)
		info.DocumentSize = len

		// Закрытие подключения
		defer resp.Body.Close()
	}

	// Переходим к следующему сервису
	service.MyService.Foo(info, c)
}