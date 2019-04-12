package main

import (
	"testing"
	"reflect"
)

func TestQuickSort(t *testing.T) {
	s:=QuickSort([]int{9, 3, -4, 6, -5, 3})
	result := []int{-5,-4, 3, 3, 6, 9}
	if reflect.DeepEqual(s, result)!=true{
		t.Errorf("Quicksort was incorrect, want %d, got %d",result,s)
	}
}

func TestPrintArray(t *testing.T) {
	s:=QuickSort([]int{9, 3, -4, 6, -5, 3})
	result := []int{-5,-4, 3, 3, 6, 9}
	if reflect.DeepEqual(s, result)!=true{
		t.Errorf("Quicksort was incorrect, want %d, got %d",result,s)
	}
}