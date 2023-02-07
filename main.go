package main

import (
	"fmt"
	"math/rand"
	"time"
)

func remove(s []int, i int) []int {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

func randomCyclicPermutation(length int) []int {
    result := make([]int, length)
    unusedIndexes := make([]int, length - 1)
    for i := 0; i < length - 1; i++ {
        unusedIndexes[i] = i + 1
    }
    currentIndex := 0
    for i := 0; i < length - 1; i++ {
        r := rand.Int() % len(unusedIndexes)
        nextInd := unusedIndexes[r]
        unusedIndexes = remove(unusedIndexes, r)
        result[currentIndex] = nextInd
        currentIndex = nextInd
    }
    return result
}

func benchmarkLatency(sizeBytes, iterations int) float64 {
	array := randomCyclicPermutation(sizeBytes / 4)
	pointer := 0
	start := time.Now()
	for i := 0; i < iterations; i++ {
		pointer = array[pointer]
	}
	return float64(time.Since(start).Nanoseconds()) / float64(iterations)
}

func main() {
	for i := 5000.0; i <= 20_000_000.0; i *= 1.2 {
		fmt.Println(int(i), " ", benchmarkLatency(int(i), 100_000_000))
	}
}
