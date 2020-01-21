package internal

import (
	"compress/gzip"
	"encoding/csv"
	"io"
	"os"
)

func UnzipRows(path string) chan []string {
	ch := make(chan []string, 10)
	go func(c chan []string) {
		f, err := os.Open(path)
		errHandle(err)
		defer f.Close()

		gr, err := gzip.NewReader(f)
		errHandle(err)
		defer gr.Close()

		cr := csv.NewReader(gr)
		for {
			rec, err := cr.Read()
			if err == io.EOF {
				break
			}
			errHandle(err)
			c <- rec
		}

		close(c)
	}(ch)
	return ch
}
