package main

import (
	"fmt"
)

type Bucket struct {
	Size int
}

type Item struct {
	Price  int
	Weight int
}

func (b *Bucket) knapSack(size int, price, weight int) {
	var item []Item
	if size > b.Size {
		fmt.Println("Your sack is not carry all this item")
		return
	}

	for i := 0; i < len(item); i++ {
		if item[i].Price+item[i+1].Price <= size {
			if item[i].Weight+item[i+1].Weight <= weight {

			}
		}
	}
}
func main() {
	fmt.Println("Test knapsack algorithm!")
}
