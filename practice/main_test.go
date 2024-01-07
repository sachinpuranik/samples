package main

import (
	"testing"
)

func Test_BubbleSort(t *testing.T) {
	input := []int{30, 15, 17, 20}
	expected := []int{15, 17, 20, 30}

	out := BubbleSort(input)
	for i := range expected {
		if out[i] != expected[i] {
			t.FailNow()
		}
	}

	input = []int{10, 20, 30}
	expected = []int{10, 20, 30}

	out = BubbleSort(input)
	for i := range expected {
		if out[i] != expected[i] {
			t.FailNow()
		}
	}

	input = []int{}
	expected = []int{}

	out = BubbleSort(input)
	if len(out) != 0 {
		t.FailNow()
	}
}

func Test_MergeSort(t *testing.T) {
	arr := []int{15, 2, 20, 17, 19, 18, 14, 3}
	expected := []int{2, 3, 14, 15, 17, 18, 19, 20}
	out := MergeSort(arr)
	if len(out) != len(expected) {
		t.Log("len mismatch")
		t.FailNow()
	}
	for i := range out {
		if out[i] != expected[i] {
			t.Log("len mismatch")
			t.FailNow()
		}
	}
}

func Test_MergeTwoSorted(t *testing.T) {
	left := []int{2, 15, 17, 20}
	right := []int{3, 14, 18, 19}
	expected := []int{2, 3, 14, 15, 17, 18, 19, 20}
	out := MergeTwoSorted(left, right)
	if len(out) != len(expected) {
		t.FailNow()
	}

	for i := range expected {
		if expected[i] != out[i] {
			t.FailNow()
		}
	}
}
