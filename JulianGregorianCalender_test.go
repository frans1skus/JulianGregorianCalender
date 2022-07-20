package main

import "testing"

func TestDayOfProgrammer_OutsideJulianGregorianCalender(t *testing.T) {
	expectation := "Year outside Julian and Gregorian Calender"
	actual := dayOfProgrammer(2701)
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestDayOfProgrammer_WithGregorianCalenderNotLeapYear(t *testing.T) {
	expectation := "13.09.2017"
	actual := dayOfProgrammer(2017)
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestDayOfProgrammer_WithGregorianCalenderLeapYear(t *testing.T) {
	expectation := "12.09.2016"
	actual := dayOfProgrammer(2016)
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestDayOfProgrammer_WithJulianCalenderNotLeapYear(t *testing.T) {
	expectation := "13.09.1801"
	actual := dayOfProgrammer(1801)
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestDayOfProgrammer_WithJulianCalenderLeapYear(t *testing.T) {
	expectation := "12.09.1800"
	actual := dayOfProgrammer(1800)
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}
