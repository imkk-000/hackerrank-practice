package main_test

import (
	. "dojo"
	"reflect"
	"testing"
)

func TestAddEmailInputFirstNameJuliaEmailJuliaAtJuliaDotMeShouldFoundRiya(t *testing.T) {
	expectedEmails := []string{}
	expectedAddStatus := false

	db := DB{}
	actualAddStatus := db.AddUser("julia", "julia@julia.me")
	actualEmails := db.GetAllEmail()

	if expectedAddStatus != actualAddStatus {
		t.Errorf("expect %v but it got %v", expectedAddStatus, actualAddStatus)
	}

	if len(expectedEmails) != len(actualEmails) {
		t.Errorf("expect %v but it got %v", expectedEmails, actualEmails)
	}
}

func TestAddEmailInputFirstNameRiyaEmailRiyaAtGmailDotComShouldFoundRiya(t *testing.T) {
	expectedEmails := []string{"riya"}
	expectedAddStatus := true

	db := DB{}
	actualAddStatus := db.AddUser("riya", "riya@gmail.com")
	actualEmails := db.GetAllEmail()

	if expectedAddStatus != actualAddStatus {
		t.Errorf("expect %v but it got %v", expectedAddStatus, actualAddStatus)
	}

	if !reflect.DeepEqual(expectedEmails, actualEmails) {
		t.Errorf("expect %v but it got %v", expectedEmails, actualEmails)
	}
}

func TestValidateEmailInputMyOwnEmailAt1mkkDotXYZShouldReturnFalse(t *testing.T) {
	expectedValidateEmail := false

	actualValidateEmail := ValidateEmail("myownemail@1mkk.xyz")

	if expectedValidateEmail != actualValidateEmail {
		t.Errorf("expect %v but it got %v", expectedValidateEmail, actualValidateEmail)
	}
}

func TestValidateEmailInputTestingEmailAtGmailDotComShouldReturnTrue(t *testing.T) {
	expectedValidateEmail := true

	actualValidateEmail := ValidateEmail("testingemail@gmail.com")

	if expectedValidateEmail != actualValidateEmail {
		t.Errorf("expect %v but it got %v", expectedValidateEmail, actualValidateEmail)
	}
}
