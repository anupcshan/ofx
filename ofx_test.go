package ofx

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func verifyOfx(t *testing.T, _ofx *Ofx, acctNum string, routingID string) {
	if _ofx == nil {
		t.Errorf("Nil ofx\n")
	}

	if _ofx.AccountNumber != acctNum {
		t.Errorf("Wrong account number. Expected: %s Actual: %s\n", acctNum, _ofx.AccountNumber)
	}

	if _ofx.RoutingCode != routingID {
		t.Errorf("Wrong routing number. Expected: %s Actual: %s\n", routingID, _ofx.RoutingCode)
	}
}

func TestParseV102(t *testing.T) {
	f, err := os.Open("testdata/v102.ofx")
	if err != nil {
		t.Fatal(err)
	}

	_ofx, err := Parse(f)
	if err != nil {
		t.Error(err)
	}

	verifyOfx(t, _ofx, "098-121", "987654321")
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
		t.Errorf("Nil ofx\n")
	}

	verifyOfx(t, _ofx, "098-121", "987654321")
}

func BenchmarkOFXParse(b *testing.B) {
	bts, err := ioutil.ReadFile("testdata/v103.ofx")
	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r := bytes.NewReader(bts)
		if _, err := Parse(r); err != nil {
			b.Errorf("Error while parsing: %v\n", err)
		}
	}
}
