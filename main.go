package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

var wg sync.WaitGroup

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	src := rand.New(rand.NewSource(time.Now().UnixNano()))
	data := make([]int, 0, size)
	for i := 0; i < size; i++ {
		randomNumber := src.Int()
		data = append(data, randomNumber)
	}
	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	maxNumber := data[0]
	for _, v := range data {
		if v > maxNumber {
			maxNumber = v
		}
	}
	return maxNumber
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	effectiveChunks := CHUNKS
	if len(data) < effectiveChunks {
		effectiveChunks = len(data)
	}
	lenChunk := len(data) / effectiveChunks
	sliceOfMax := make([]int, effectiveChunks)
	for i := 0; i < effectiveChunks; i++ {
		startChunk := i * lenChunk
		endChunk := startChunk + lenChunk
		if i == CHUNKS {
			endChunk = len(data)
		}
		chunk := data[startChunk:endChunk]

		wg.Add(1)
		go func(chunk []int, i int) {
			defer wg.Done()
			sliceOfMax[i] = maximum(chunk)
		}(chunk, i)
	}
	wg.Wait()
	max := maximum(sliceOfMax)
	return max
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	timeStart := time.Now()
	max := maximum(data)
	elapsed := time.Since(timeStart).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	timeStart = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(timeStart).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
