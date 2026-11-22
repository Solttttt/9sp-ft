package main

import (
	"testing"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	sizeTest := 10_000
	dataTest := generateRandomElements(sizeTest)
	if len(dataTest) != sizeTest {
		t.Errorf("slice empty")
	}

	sizeTest = 0
	dataTest = generateRandomElements(sizeTest)
	if len(dataTest) != 0 {
		t.Error("expected empty slice")
	}

	sizeTest = 10_000
	dataTest = generateRandomElements(sizeTest)
	for _, v := range dataTest {
		if v < 0 {
			t.Errorf("negative numbers in slice")
			return
		}
	}
}

func TestMaximum(t *testing.T) {
	if maximum([]int{}) != 0 {
		t.Error("empty slice")
	}

	dataTest := []int{1, 1, 1, 1, 1, 1}
	maxTest := maximum(dataTest)
	if maxTest != dataTest[0] {
		t.Errorf("not correct answer")
	}

	dataTest = []int{1}
	maxTest = maximum(dataTest)
	if maxTest != dataTest[0] {
		t.Error("slice consit of one element")
	}

	sizeTest := 10_000
	dataTest = generateRandomElements(sizeTest)
	maxTest = maximum(dataTest)
	for _, v := range dataTest {
		if v > maxTest {
			t.Errorf("found element haier than  max")
			return
		}
	}
}

func TestMaxChunks(t *testing.T) {
	if maxChunks([]int{}) != 0 {
		t.Errorf("empty slice")
	}

	dataTest := []int{1, 2, 3, 4, 5}
	if maximum(dataTest) != maxChunks(dataTest) {
		t.Error("chunk smoller")
	}

	sizeTest := 99_000_000
	dataTest = generateRandomElements(sizeTest)
	if maximum(dataTest) != maxChunks(dataTest) {
		t.Error("chunk equal")
	}

	sizeTest = 100_000_000
	dataTest = generateRandomElements(sizeTest)
	if maximum(dataTest) != maxChunks(dataTest) {
		t.Error("chunk equal")
	}

	sizeTest = 100_000_001
	dataTest = generateRandomElements(sizeTest)
	if maximum(dataTest) != maxChunks(dataTest) {
		t.Error("chunk large")
	}
}
