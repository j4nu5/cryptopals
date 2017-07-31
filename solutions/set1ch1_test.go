package solutions

import (
	"testing"
)

func TestHexToBase64(t *testing.T) {
	in := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	out, err := HexToBase64(in)
	if err != nil {
		t.Error("For", in, "unexpected", err)
	}

	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	if out != expected {
		t.Error("For", in, "expected", expected, "got", out)
	}
}
