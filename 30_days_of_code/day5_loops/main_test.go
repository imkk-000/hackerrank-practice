package main_test

import (
	dojo "dojo"
	"reflect"
	"testing"
)

func TestGetMultiplicationTableInput2ShouldReturnMultiplicationTableOf2(t *testing.T) {
	expectedMultiplicationTable := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	actualMultiplicationTable := dojo.GetMultiplicationTable(2)

	if !reflect.DeepEqual(expectedMultiplicationTable, actualMultiplicationTable) {
		t.Errorf("expect %v but it got %v", expectedMultiplicationTable, actualMultiplicationTable)
	}
}
