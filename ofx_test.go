package ofx

import (
	"os"
	"testing"
)

func TestParseV102(t *testing.T) {
	f, err := os.Open("testdata/v102.ofx")
	if err != nil {
		t.Fatal(err)
	}

	_ofx, err := Parse(f)
	if err != nil {
		t.Error(err)
	}

	if _ofx == nil {
		t.Errorf("Nil ofx %s\n", _ofx)
	}
}

func TestParseV103(t *testing.T) {
	f, err := os.Open("testdata/v103.ofx")
	if err != nil {
		t.Fatal(err)
	}

	_ofx, err := Parse(f)
	if err != nil {
		t.Error(err)
	}

	if _ofx == nil {
		t.Errorf("Nil ofx %s\n", _ofx)
	}
}
