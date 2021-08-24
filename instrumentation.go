package main

import "sync"

// Сервис для сбора статистики нагрузки процессов
type statisticService struct {
	data map[int]int
	sync.RWMutex
	MyService
}

func (service statisticService) Foo(info ResponseInfo, c chan int) {
	// Увеличиваем счетчик запросов ядра
	service.Lock()
	if _, ok := service.data[info.ProcNum]; !ok{
		service.data[info.ProcNum] = 1
	}else{
		service.data[info.ProcNum]++
	}
	service.Unlock()

	// Заполняем канал для освобождения ядра
	defer func() {
		c <- 1
	}()

	// Переходим к следующему сервису
	service.MyService.Foo(info, c)
}
