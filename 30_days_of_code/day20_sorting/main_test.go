package main_test

import (
	dojo "dojo"
	"reflect"
	"testing"
)

func TestBubbleSortWithSwapTimeInputSlice3And2And1ShoudReturnSortedSliceAndSwap1Time(t *testing.T) {
	expectedSortedSlice, expectedSwapTime := []int{1, 2, 3}, 1

	actualSortedSlice, actualSwapTime := dojo.BubbleSortWithSwapTime([]int{1, 3, 2})

	if !reflect.DeepEqual(expectedSortedSlice, actualSortedSlice) {
		t.Errorf("expect %v but it got %v", expectedSortedSlice, actualSortedSlice)
	}

	if expectedSwapTime != actualSwapTime {
		t.Errorf("expect %d but it got %d", expectedSwapTime, actualSwapTime)
	}
}

func TestBubbleSortWithSwapTimeInputSlice3And2And1ShoudReturnSortedSliceAndSwap0Time(t *testing.T) {
	expectedSortedSlice, expectedSwapTime := []int{1, 2, 3}, 0

	actualSortedSlice, actualSwapTime := dojo.BubbleSortWithSwapTime([]int{1, 2, 3})

	if !reflect.DeepEqual(expectedSortedSlice, actualSortedSlice) {
		t.Errorf("expect %v but it got %v", expectedSortedSlice, actualSortedSlice)
	}

	if expectedSwapTime != actualSwapTime {
		t.Errorf("expect %d but it got %d", expectedSwapTime, actualSwapTime)
	}
}

func TestBubbleSortWithSwapTimeInputSlice3And2And1ShoudReturnSortedSliceAndSwap3Times(t *testing.T) {
	expectedSortedSlice, expectedSwapTime := []int{1, 2, 3}, 3

	actualSortedSlice, actualSwapTime := dojo.BubbleSortWithSwapTime([]int{3, 2, 1})

	if !reflect.DeepEqual(expectedSortedSlice, actualSortedSlice) {
		t.Errorf("expect %v but it got %v", expectedSortedSlice, actualSortedSlice)
	}

	if expectedSwapTime != actualSwapTime {
		t.Errorf("expect %d but it got %d", expectedSwapTime, actualSwapTime)
	}
}
