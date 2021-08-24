package main

import (
	"bufio"
	"fmt"
	"github.com/xlab/closer"
	"os"
	"runtime"
)

func main() {
	// Устанавливаем максимальное кол-во параллельных процессов
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Установка сервисов и их очередности
	var service MyService
	service = outputService{}
	service = clientService{MyService: service}
	service = statisticService{data: make(map[int]int), MyService: service}

	// Обработчик завершения программы (CTRL+C)
	closer.Bind(func() {
		os.Stdout.WriteString("\nСтатистика\n")

		data := service.(statisticService).data

		for key, value := range data{
			os.Stdout.WriteString(fmt.Sprintf("Поток %d: %d\n", key, value))
		}

		os.Stdout.WriteString("Конец\n")
	})

	// Чтение данных со стандартного входа
	Foo(service)

	// Блокировка в ожидании сигнала завершения программы
	closer.Hold()
}

// Основная функция
func Foo(service MyService) {
	// Кол-во парных ядер
	var numCPU = runtime.NumCPU()

	// Канал для блокировки
	c := make(chan int, numCPU)

	for {
		// Распределение запросов по потокам
		for i := 0; i < numCPU; i++ {
			url := ReadLine()

			go service.Foo(ResponseInfo{URL: url, ProcNum: i}, c)
		}

		// Ожидание освобождения ядер
		for i := 0; i < numCPU; i++ {
			<-c
		}
	}
}

// Чтение строки с стандартного входа
func ReadLine() (url string) {
	// Стандартный ввод
	stdin := bufio.NewReader(os.Stdin)

	// Чтение строки
	line, err := stdin.ReadBytes('\n')

	if err != nil {
		os.Stdout.WriteString(err.Error())
		return url
	}

	// Удаление символа \n в конце строки
	url = string(line[:len(line)-1])

	return url

}
