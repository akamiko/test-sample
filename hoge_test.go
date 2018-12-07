package hoge

import "testing"

func TestHoge(t *testing.T) {
	actual := hoge()
	expected := "hoge"
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}
