package main

import "testing"

var (

	num Calculator = Calculator{2,3}

	hasilPenambahan int = 5

	hasilPengurangan int = -1

)


func TestCalculator_Add(t *testing.T){
	t.Logf("Hasil Penjumlahan : %d", num.Add())

	if num.Add() != hasilPenambahan {
		t.Errorf("Salah! harusnya %d", hasilPenambahan)
	}
}


func TestCalcaultor_Sub(t *testing.T){
	t.Logf("Hasil Pengurangan : %d", num.Sub())

	if num.Sub() != hasilPengurangan {
		t.Errorf("Salah! harusnya %d", hasilPengurangan)
	}

}