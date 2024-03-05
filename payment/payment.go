package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	p := Payment{
		From:   "anon",
		To:     "sara",
		Amount: 15,
	}

	p.Process()
	p.Process()
}

type Payment struct {
	From   string
	To     string
	Amount float64

	once sync.Once
}

func (payment *Payment) Process() {
	ts := time.Now()
	payment.once.Do(func() {
		payment.processTransaction(ts)
	})
}

func (payment *Payment) processTransaction(t time.Time) {
	ts := t.Format(time.RFC3339)
	fmt.Printf("[%s] %s -> $%.2f -> %s \n", ts, payment.From, payment.Amount, payment.To)
}
