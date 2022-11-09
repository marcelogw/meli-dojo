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

func (c *Char) IncrementKi(ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			c.Power++
			if c.Power >= 8000 {
				return
			}
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	chars := []*Char{
		{Name: "Goku"},
		{Name: "Gohan"},
		{Name: "Vegeta"},
	}
	for _, v := range chars {
		go v.IncrementKi(ctx, cancel)
	}
	<-ctx.Done()

	printChars(chars)
	time.Sleep(time.Second) // ensure no one is still increasing ki
	printChars(chars)
}
