package main

import (
	"testing"
)

func TestPrintIAmOldInputAge18ShoudlReturnYouAreOldMessage(t *testing.T) {
	expectedMessage := "You are old."

	actualMessage := printIAmOld(18)

	if expectedMessage != actualMessage {
		t.Errorf("expect %s but it got %s", expectedMessage, actualMessage)
	}
}

func TestPrintIAmOldInputAge13ShouldReturnYouAreATeenagerMessage(t *testing.T) {
	expectedMessage := "You are a teenager."

	actualMessage := printIAmOld(13)

	if expectedMessage != actualMessage {
		t.Errorf("expect %s but it got %s", expectedMessage, actualMessage)
	}
}

func TestPrintIAmOldInputAge12ShouldReturnYouAreYoungMessage(t *testing.T) {
	expectedMessage := "You are young."

	actualMessage := printIAmOld(12)

	if expectedMessage != actualMessage {
		t.Errorf("expect %s but it got %s", expectedMessage, actualMessage)
	}
}

func TestNewPersonInputAgeMinus1ShouldBeAge0(t *testing.T) {
	expectedAgeOfPerson := 0

	p := person{}
	p.NewPerson(-1)
	actualAgeOfPerson := p.age

	if expectedAgeOfPerson != actualAgeOfPerson {
		t.Errorf("expect %d but it got %d", expectedAgeOfPerson, actualAgeOfPerson)
	}
}

func TestNewPersonInputAge5ShouldBeAge5(t *testing.T) {
	expectedAgeOfPerson := 5

	p := person{}
	p = p.NewPerson(5)
	actualAgeOfPerson := p.age

	if expectedAgeOfPerson != actualAgeOfPerson {
		t.Errorf("expect %d but it got %d", expectedAgeOfPerson, actualAgeOfPerson)
	}
}