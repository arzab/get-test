package main

import "time"

//ResponseInfo - структура для сбора информации по результатам запроса
type ResponseInfo struct {
	URL          string
	Duration     time.Duration
	StatusCode   int
	DocumentSize int
	Err          string
	ProcNum		 int
}
