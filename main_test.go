package main

import "testing"

func TestMostCommonWords_001_TextMoreWordsThanQuantity(t *testing.T) {
	body := "one two two three three three four four four four five five five % (aa @) five 1 five"
	result := MostCommonWords(body, 4)
	if len(result) != 4 {
		t.Errorf("want 4, got %d", len(result))
	}
	if result["five"] != 5 {
		t.Errorf("want 5, got %d", result["five"])
	}
	if result["four"] != 4 {
		t.Errorf("want 4, got %d", result["four"])
	}
	if result["three"] != 3 {
		t.Errorf("want 3, got %d", result["three"])
	}
	if result["two"] != 2 {
		t.Errorf("want 2, got %d", result["two"])
	}
}

func TestMostCommonWords_002_TextLessWordsThanQuantity(t *testing.T) {
	body := "one two two three three three a-a"
	result := MostCommonWords(body, 10)
	if len(result) != 4 {
		t.Errorf("want 4, got %d", len(result))
	}
}

func TestMostCommonWords_002_NonLettersCharacters(t *testing.T) {
	body := "234 #### & ()"
	result := MostCommonWords(body, 10)
	if len(result) != 0 {
		t.Errorf("want 0, got %d", len(result))
	}
}