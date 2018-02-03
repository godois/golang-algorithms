package main

import (
	"fmt"
	"math/rand"
)

const arraySize = 1000000

func main() {
	var numbersToOrder = prepareArray()
	fmt.Println("Ordering numbers... ")
	var outPut = applySorting(numbersToOrder)
	fmt.Println("Result: ", outPut)
}

func prepareArray() [arraySize]int {

	var array = [arraySize]int{}
	for i := 0; i < arraySize; i++ {
		array[i] = rand.Intn(arraySize)
	}
	return array
}

func applySorting(numbersToOrder [arraySize]int) [arraySize]int {

	for j := 0; j < len(numbersToOrder); j++ {
		currentValue := numbersToOrder[j]
		i := j - 1
		for i >= 0 && numbersToOrder[i] > currentValue {
			numbersToOrder[i+1] = numbersToOrder[i]
			i = i - 1
		}
		numbersToOrder[i+1] = currentValue
	}

	return numbersToOrder
}
