package main_test

import (
	dojo "dojo"
	"reflect"
	"testing"
)

func TestGetContactInputNameEdwardShouldReturnStringNotFound(t *testing.T) {
	expectedContactInformation := "Not found"

	contactBook := dojo.ContactBook{}
	contactBook = contactBook.AddContact("sam", "12345")
	actualContactInformation := contactBook.GetContact("edward")

	if expectedContactInformation != actualContactInformation {
		t.Errorf("expect '%s' but it got '%s'", expectedContactInformation, actualContactInformation)
	}
}

func TestGetContactInputNameSamShouldReturnStringSamEqual12345(t *testing.T) {
	expectedContactInformation := "sam=12345"

	contactBook := dojo.ContactBook{}
	contactBook = contactBook.AddContact("sam", "12345")
	actualContactInformation := contactBook.GetContact("sam")

	if expectedContactInformation != actualContactInformation {
		t.Errorf("expect '%s' but it got '%s'", expectedContactInformation, actualContactInformation)
	}
}

func TestAddContactInputNameSamTelephoneNumber12345ShouldReturnContactBook(t *testing.T) {
	expectedContactBook := dojo.ContactBook{
		"sam": "12345",
	}

	contactBook := dojo.ContactBook{}
	actualContactBook := contactBook.AddContact("sam", "12345")

	if !reflect.DeepEqual(expectedContactBook, actualContactBook) {
		t.Errorf("expect %v but it got %v", expectedContactBook, actualContactBook)
	}
}
