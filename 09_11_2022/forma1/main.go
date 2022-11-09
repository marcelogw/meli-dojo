package main

import (
	"context"
	"fmt"
	"time"
)

type Char struct {
	Name  string
	Power int
}

func printChars(c []*Char) {
	fmt.Printf("Final result:\n")
	for _, v := range c {
		fmt.Printf("%s: %d\n", v.Name, v.Power)
	}
	fmt.Println("---")
}

func (c *Char) IncrementKi(ctx context.Context, ch chan bool) {
	for {
		if ctx.Err() == nil {
			c.Power++
			if c.Power >= 8000 {
				ch <- true
				return
			}
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan bool) // use channel to control the return of some goroutine

	chars := []*Char{
		{Name: "Goku"},
		{Name: "Gohan"},
		{Name: "Vegeta"},
	}
	for _, v := range chars {
		go v.IncrementKi(ctx, ch)
	}

	<-ch
	cancel()

	printChars(chars)
	time.Sleep(time.Second) // ensure no one is still increasing ki
	printChars(chars)
}
