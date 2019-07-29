package main_test

import (
	. "dojo"
	"testing"
)

func TestCalculateFineInputActualDateD2M1Y2015ExpectedDateD1M1Y2016ShoudlReturnFine0(t *testing.T) {
	expectedFine := 0
	inputActualDate := Date{Day: 2, Month: 1, Year: 2015}
	inputExpectedDate := Date{Day: 1, Month: 1, Year: 2016}

	actualFine := CalculateFine(inputActualDate, inputExpectedDate)

	if expectedFine != actualFine {
		t.Errorf("expect %d but it got %d", expectedFine, actualFine)
	}
}

func TestCalculateFineInputActualDateD1M2Y2015ExpectedDateD1M1Y2016ShoudlReturnFine0(t *testing.T) {
	expectedFine := 0
	inputActualDate := Date{Day: 1, Month: 2, Year: 2015}
	inputExpectedDate := Date{Day: 1, Month: 1, Year: 2016}

	actualFine := CalculateFine(inputActualDate, inputExpectedDate)

	if expectedFine != actualFine {
		t.Errorf("expect %d but it got %d", expectedFine, actualFine)
	}
}

func TestCalculateFineInputActualDateD1M1Y2015ExpectedDateD1M1Y2016ShoudlReturnFine0(t *testing.T) {
	expectedFine := 0
	inputActualDate := Date{Day: 1, Month: 1, Year: 2015}
	inputExpectedDate := Date{Day: 1, Month: 1, Year: 2016}

	actualFine := CalculateFine(inputActualDate, inputExpectedDate)

	if expectedFine != actualFine {
		t.Errorf("expect %d but it got %d", expectedFine, actualFine)
	}
}

func TestCalculateFineInputActualDateD1M1Y2015ExpectedDateD1M2Y2015ShoudlReturnFine0(t *testing.T) {
	expectedFine := 0
	inputActualDate := Date{Day: 1, Month: 1, Year: 2015}
	inputExpectedDate := Date{Day: 1, Month: 2, Year: 2015}

	actualFine := CalculateFine(inputActualDate, inputExpectedDate)

	if expectedFine != actualFine {
		t.Errorf("expect %d but it got %d", expectedFine, actualFine)
	}
}

func TestCalculateFineInputActualDateD1M1Y2015ExpectedDateD2M1Y2015ShoudlReturnFine0(t *testing.T) {
	expectedFine := 0
	inputActualDate := Date{Day: 1, Month: 1, Year: 2015}
	inputExpectedDate := Date{Day: 2, Month: 1, Year: 2015}

	actualFine := CalculateFine(inputActualDate, inputExpectedDate)

	if expectedFine != actualFine {
		t.Errorf("expect %d but it got %d", expectedFine, actualFine)
	}
}

func TestCalculateFineInputActualDateD1M1Y2015ExpectedDateD1M1Y2015ShoudlReturnFine0(t *testing.T) {
	expectedFine := 0
	inputActualDate := Date{Day: 1, Month: 1, Year: 2015}
	inputExpectedDate := Date{Day: 1, Month: 1, Year: 2015}

	actualFine := CalculateFine(inputActualDate, inputExpectedDate)

	if expectedFine != actualFine {
		t.Errorf("expect %d but it got %d", expectedFine, actualFine)
	}
}

func TestCalculateFineInputActualDateD6M6Y2016ExpectedDateD6M6Y2015ShoudlReturnFine10000(t *testing.T) {
	expectedFine := 10000
	inputActualDate := Date{Day: 6, Month: 6, Year: 2016}
	inputExpectedDate := Date{Day: 6, Month: 6, Year: 2015}

	actualFine := CalculateFine(inputActualDate, inputExpectedDate)

	if expectedFine != actualFine {
		t.Errorf("expect %d but it got %d", expectedFine, actualFine)
	}
}

func TestCalculateFineInputActualDateD6M7Y2015ExpectedDateD6M6Y2015ShoudlReturnFine500(t *testing.T) {
	expectedFine := 500
	inputActualDate := Date{Day: 6, Month: 7, Year: 2015}
	inputExpectedDate := Date{Day: 6, Month: 6, Year: 2015}

	actualFine := CalculateFine(inputActualDate, inputExpectedDate)

	if expectedFine != actualFine {
		t.Errorf("expect %d but it got %d", expectedFine, actualFine)
	}
}

func TestCalculateFineInputActualDateD9M6Y2015ExpectedDateD6M6Y2015ShoudlReturnFine45(t *testing.T) {
	expectedFine := 45
	inputActualDate := Date{Day: 9, Month: 6, Year: 2015}
	inputExpectedDate := Date{Day: 6, Month: 6, Year: 2015}

	actualFine := CalculateFine(inputActualDate, inputExpectedDate)

	if expectedFine != actualFine {
		t.Errorf("expect %d but it got %d", expectedFine, actualFine)
	}
}

func TestGetDateInput22Space11Space2015ShouldReturnStructDateD22M11Y2015(t *testing.T) {
	expectedDate := Date{
		Day:   22,
		Month: 11,
		Year:  2015,
	}

	actualDate := GetDate("22 11 2015")

	if expectedDate != actualDate {
		t.Errorf("expect %d but it got %d", expectedDate, actualDate)
	}
}

func TestGetDateInput9Space6Space2015ShouldReturnStructDateD9M6Y2015(t *testing.T) {
	expectedDate := Date{
		Day:   9,
		Month: 6,
		Year:  2015,
	}

	actualDate := GetDate("9 6 2015")

	if expectedDate != actualDate {
		t.Errorf("expect %d but it got %d", expectedDate, actualDate)
	}
}
