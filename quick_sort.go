package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Función que implementa el algoritmo QuickSort para ordenar un arreglo
func quickSort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high)
		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high)
	}
}

// Función que realiza la partición del arreglo para QuickSort
func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}

// Función que calcula la mediana de un arreglo ordenado
func calculateMedian(arr []int) float64 {
	n := len(arr)
	if n%2 == 0 {
		return float64(arr[n/2-1]+arr[n/2]) / 2.0
	} else {
		return float64(arr[n/2])
	}
}

func main() {
	size := 1000001
	arr := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(1000000)
	}

	start := time.Now()
	quickSort(arr, 0, len(arr)-1)
	median := calculateMedian(arr)
	elapsed := time.Since(start)

	fmt.Printf("La mediana es: %.2f\n", median)
	fmt.Printf("Tiempo de ejecución: %s\n", elapsed)
}
