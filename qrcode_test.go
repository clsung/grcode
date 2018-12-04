package grcode

import (
	"testing"
)

var testImgFile = "data/test.png"

func TestGetDataFromFile(t *testing.T) {
	results, err := GetDataFromFile(testImgFile)
	if err != nil {
		t.Fatal(err)
	}
	if want, got := 1, len(results); want != got {
		t.Errorf("wrong request %q, want %q", got, want)
	}
	if want, got := "https://github.com/clsung/grcode", results[0]; want != got {
		t.Errorf("wrong request %q, want %q", got, want)
	}
}
