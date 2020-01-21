package main

import (
	"enrich/internal"
	"fmt"
)

func main() {
	rows := internal.UnzipRows("/home/alexander/mygoprojects/enrich/cmd/gz/test_data/test.csv.gz")

	for r := range rows {
		fmt.Println(r)
	}
}
