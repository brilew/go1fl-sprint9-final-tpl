package main

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestMaximumEmpty(t *testing.T) {
	assert.Equal(t, maximum(nil), 0)
	assert.Equal(t, maximum([]int{}), 0)
}
func TestMaximumPositive(t *testing.T) {
	assert.Equal(t, 9, maximum([]int{1, 5, 9, 3, 7}))
}
func TestMaximumSingle(t *testing.T) {
	assert.Equal(t, 42, maximum([]int{42}))
}
func TestMaximumAtStart(t *testing.T) {
	assert.Equal(t, 100, maximum([]int{100, 5, 3, 1}))
}

func TestMaximumAtEnd(t *testing.T) {
	assert.Equal(t, 100, maximum([]int{1, 3, 5, 100}))
}

func TestMaximumAllEqual(t *testing.T) {
	assert.Equal(t, 7, maximum([]int{7, 7, 7, 7}))
}
func TestMaximumNegative(t *testing.T) {
	assert.Equal(t, maximum([]int{-5, -3, -10}), -3)
}

func TestMaximumMixed(t *testing.T) {
	assert.Equal(t, 5, maximum([]int{-10, 5, -3, 0, 2}))
}

// теперь тесты функции с горутинами
func TestMaxChunksNil(t *testing.T) {
	assert.Equal(t, 0, maxChunks(nil))
}

func TestMaxChunksEmpty(t *testing.T) {
	assert.Equal(t, 0, maxChunks([]int{}))
}

func TestMaxChunksSingle(t *testing.T) {
	assert.Equal(t, 42, maxChunks([]int{42}))
}

func TestMaxChunksSmall(t *testing.T) {
	assert.Equal(t, 5, maxChunks([]int{1, 5, 3, 2}))
}

func TestMaxChunksKnownMaxInMiddle(t *testing.T) {
	data := make([]int, CHUNKS*10)
	for i := range data {
		data[i] = i
	}
	data[len(data)/2] = 1000000
	assert.Equal(t, 1000000, maxChunks(data))
}

func TestMaxChunksKnownMaxInLastChunk(t *testing.T) {
	// максимум именно в последнем чанке, чтобы проверить обработку остатка
	data := make([]int, CHUNKS*5+3)
	for i := range data {
		data[i] = i
	}
	data[len(data)-1] = 999999
	assert.Equal(t, 999999, maxChunks(data))
}

func TestMaxChunksKnownMaxInFirstChunk(t *testing.T) {
	data := make([]int, CHUNKS*5)
	for i := range data {
		data[i] = i
	}
	data[0] = 1000000
	assert.Equal(t, 1000000, maxChunks(data))
}

func TestMaxChunksAllEqual(t *testing.T) {
	data := make([]int, CHUNKS*5)
	for i := range data {
		data[i] = 7
	}
	assert.Equal(t, 7, maxChunks(data))
}
