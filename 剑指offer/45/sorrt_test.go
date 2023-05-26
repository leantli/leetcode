package main

import (
	"math/rand"
	"testing"
)

func BenchmarkQuickSort(b *testing.B) {
	arr := make([]int, 0, 10000)
	for i := 0; i < 10000; i++ {
		arr = append(arr, rand.Int())
	}
	for n := 0; n < b.N; n++ {
		quickSort(arr, 0, len(arr)-1)
	}
}

func BenchmarkQuickSort2(b *testing.B) {
	arr := make([]int, 0, 10000)
	for i := 0; i < 10000; i++ {
		arr = append(arr, rand.Int())
	}
	for n := 0; n < b.N; n++ {
		quickSort2(arr, 0, len(arr)-1)
	}
}
