package main

import (
	"fmt"
	"log"
)

func safeDiv(a, b int) (q int, err error) {
	defer func() {
		if e := recover(); e != nil {
			log.Println("ERROR:", e)
			err = fmt.Errorf("%v", e)
		}
	}()

	return a / b, nil
}

func main() {
	fmt.Println(safeDiv(1, 0))
}
