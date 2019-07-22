package main

import (
	"reflect"
	"testing"
)

func TestReverseListOfNumbersInputOddListShouldReturnReverseOfList(t *testing.T) {
	expectedList := []int{2, 3, 4, 1, 5}

	actualList := ReverseListOfNumbers([]int{5, 1, 4, 3, 2})

	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("expect %v but it got %v", expectedList, actualList)
	}
}

func TestReverseListOfNumbersInputEvenListShouldReturnReverseOfList(t *testing.T) {
	expectedList := []int{2, 3, 4, 1}

	actualList := ReverseListOfNumbers([]int{1, 4, 3, 2})

	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("expect %v but it got %v", expectedList, actualList)
	}
}
