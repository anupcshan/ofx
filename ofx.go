package ofx

import (
	"bytes"
	"encoding/xml"
	"io"
	"log"
	"strings"
)

// AccountType indicates type of account represented by OFX document.
type AccountType int

const (
	// UNKNOWN - Account type could not be determined
	UNKNOWN AccountType = iota
	// CHECKING - Checking account
	CHECKING AccountType = iota
	// SAVING - Savings account
	SAVING AccountType = iota
)

type nextKey int

const (
	none      nextKey = iota
	acctID    nextKey = iota
	routingID nextKey = iota
)

// Ofx contains a parsed Ofx document.
type Ofx struct {
	ActType       AccountType
	RoutingCode   string
	AccountNumber string
}

// Parse parses an input stream and produces an Ofx instance summarizing it. In case of any errors
// during the parse, a non-nil error is returned.
func Parse(f io.Reader) (*Ofx, error) {
	ofx := &Ofx{}
	stack := make([]string, 1000)
	stackPos := 0

	next := none

	dec := xml.NewDecoder(f)

	tok, err := dec.RawToken()
	for err == nil {
		switch t := tok.(type) {
		case xml.StartElement:
			stack[stackPos] = t.Name.Local
			stackPos++

			switch t.Name.Local {
			case "ACCTID":
				next = acctID

			case "BANKID":
				next = routingID
			}

		case xml.CharData:
			var b bytes.Buffer
			if _, err := b.Write(t); err != nil {
				return nil, err
			}
			res := strings.TrimSpace(b.String())

			switch next {
			case acctID:
				ofx.AccountNumber = res

			case routingID:
				ofx.RoutingCode = res
			}

			next = none

		case xml.EndElement:
			for stackPos != 0 {
				if stack[stackPos-1] == t.Name.Local {
					stackPos--
					break
				}
				stackPos--
			}

		default:
			log.Printf("Unknown: %T %s\n", t, t)
		}

		tok, err = dec.RawToken()

		if err != nil && err != io.EOF {
			log.Printf("Error: %s\n", err)
		}
	}

	return ofx, nil
}
