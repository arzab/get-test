package main

import (
	"fmt"
	"os"
)

// Сервис для вывода результата запроса
type outputService struct {
	MyService
}

// Вывод данных в формате CSV
func (outputService) Foo(info ResponseInfo, c chan int) {
	var output string

	if info.Err == "" {
		output = fmt.Sprintf("%s;%d;%d;%dms", info.URL, info.StatusCode, info.DocumentSize, info.Duration.Milliseconds())
	} else {
		output = fmt.Sprintf("Url:%s\nError:%s\n", info.URL, info.Err)
	}

	os.Stdout.WriteString(output)
}

