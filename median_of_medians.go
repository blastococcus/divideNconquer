package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Función que implementa el algoritmo Median of Medians
func medianOfMedians(arr []int, k int) int {
	if len(arr) <= 5 {
		sort.Ints(arr)
		return arr[k-1]
	}

	for i := 0; i < len(arr); i += 5 {
		end := i + 5
		if end > len(arr) {
			end = len(arr)
		}
		sort.Ints(arr[i:end])
	}

	medians := make([]int, 0, (len(arr)+4)/5)
	for i := 2; i < len(arr); i += 5 {
		medians = append(medians, arr[i])
	}

	median := medianOfMedians(medians, len(medians)/2+1)

	i := 0
	for j := 0; j < len(arr); j++ {
		if arr[j] < median {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	if k <= i {
		return medianOfMedians(arr[:i], k)
	} else if k > i+1 {
		return medianOfMedians(arr[i+1:], k-i-1)
	}
	return median
}

func main() {
	size := 1000001
	arr := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(1000000)
	}

	k := (size / 2) + 1
	start := time.Now()
	median := medianOfMedians(arr, k)
	elapsed := time.Since(start)

	fmt.Printf("La mediana es: %d\n", median)
	fmt.Printf("Tiempo de ejecución: %s\n", elapsed)
}
