package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Función que implementa la selección aleatoria para encontrar el k-ésimo elemento más pequeño
func randomizedSelect(arr []int, k int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	rand.Seed(time.Now().UnixNano())
	pivotIndex := rand.Intn(len(arr))
	pivot := arr[pivotIndex]

	low := []int{}
	high := []int{}
	equal := []int{}

	for _, num := range arr {
		if num < pivot {
			low = append(low, num)
		} else if num > pivot {
			high = append(high, num)
		} else {
			equal = append(equal, num)
		}
	}

	if k <= len(low) {
		return randomizedSelect(low, k)
	} else if k <= len(low)+len(equal) {
		return pivot
	} else {
		return randomizedSelect(high, k-len(low)-len(equal))
	}
}

func main() {
	size := 100001
	arr := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(100000)
	}

	k := (size / 2) + 1
	start := time.Now()
	median := randomizedSelect(arr, k)
	elapsed := time.Since(start)

	fmt.Printf("La mediana es: %d\n", median)
	fmt.Printf("Tiempo de ejecución: %s\n", elapsed)
}
