package main

import "testing"
func TestBinarySearch(t *testing.T) {
	var numbas = []int{1, 2, 3, 5, 8, 13, 21, 5555};

	target := 8;
	result := BinarySearch(numbas, target);
	expected := 4

	if result != expected {
        t.Errorf("binarySearch(%v, %d) should have returned index %d", numbas, target, expected);
	}

	target = 21;
	result= BinarySearch(numbas, target);
	expected = 6;

	if result != expected {	
        t.Errorf("binarySearch(%v, %d) should have returned index %d", numbas, target, expected);
	}

	target = 14;
	result = BinarySearch(numbas, target);
	expected = -1;

	if result != expected {	
        t.Errorf("binarySearch(%v, %d) should have returned index %d", numbas, target, expected);
	}

}