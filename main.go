package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}
	slice := make([]int, 0, size)
	for i := 0; i < size; i++ {
		v := rand.IntN(size)
		slice = append(slice, v)

	}

	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	var max int = data[0]
	for _, num := range data {
		if num > max {
			max = num
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if data == nil {
		return 0
	}
	var max int
	var chankSize int
	// Проверка что чанки не равны 0 или 1. Это константа. Но кто знает, что будет дальше??
	if CHUNKS <= 1 {
		max = maximum(data)
		return max
	}
	// Проверка что количество чанков не больше количества чисел в слайсе.
	// Если чисел меньше то просто находим максимум слайса.
	if len(data) <= CHUNKS {
		max = maximum(data)
		return max
	}

	chankSize = len(data) / CHUNKS
	results := make(chan int, CHUNKS)
	var wg sync.WaitGroup

	// запускаем цикд столько раз сколько чанков, что бы пройти их все
	for j := 0; j < CHUNKS; j++ {

		start := j * chankSize
		end := start + chankSize
		if j == CHUNKS-1 {
			end = len(data)
		}
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			results <- maximum(data[start:end])

		}(start, end)
	}

	wg.Wait()
	close(results)
	for m := range results {
		if m > max {
			max = m
		}
	}

	return max
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// в задании написано - Время нахождения максимумов измеряйте в Микросекундах.
	// они обозначаются как µs.
	// менять ms на µs не стал
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
