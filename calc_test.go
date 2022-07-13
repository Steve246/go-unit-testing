package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculator_Add(t *testing.T) {

	myCalc := Calculator{Num1:4, Num2:4}

	actual := myCalc.Add()

	expected := 8

	assert.Equal(t, actual, expected , "it should return 8")

}

func TestCalculator_Sub(t *testing.T) {

	myCalc := Calculator{Num1:5, Num2:4}

	actual := myCalc.Sub()

	expected := 1

	assert.Equal(t, actual, expected , "it should return 1")

}

func TestMainApp(t *testing.T) {
	os.Args = []string{"-num1", "2", "-num2", "3", "-opr", "add"}
	main()

}

