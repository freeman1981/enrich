package main

import (
	"flag"
	"fmt"
)

var testParam string

func init() {
	flag.StringVar(&testParam, "param", "value foo", "usage")
}

func main() {
	flag.Parse()
	fmt.Println(testParam)

	// грайсфул шатдаун
	// слушаем рэббит
	// прогреваем кэш (работа с постгресом)
	// распаковываем gzip пишем в click house
}
