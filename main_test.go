package main

import "testing"

func TestContains(t *testing.T){
	slice := []string{"a", "b", "c"}
	slice1 := []string{"Hello", "hi", "okay"}
	got1 := contains(slice, "a")
	got2 := contains(slice1, "b")

	want1 := true
	want2 := false

	if got1 != want1 {
		t.Error( got1, want1)
	}

	if got2 != want2 {
		t.Error( got2, want2)
	}
}

// main_test.go
func TestRemoveElementFromSlice(t *testing.T){
	arr := []string{"a", "b", "c", "d"}
	el := "a"

	a := removeElementFromSlice(arr, el)
	include := contains(a, el)

	if include {
		t.Error("The function is not working as expected!")
	}
}