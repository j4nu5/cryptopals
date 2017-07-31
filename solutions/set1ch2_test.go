package solutions

import (
	"testing"
)

func TestXorHexStrings(t *testing.T) {
	a := "1c0111001f010100061a024b53535009181c"
	b := "686974207468652062756c6c277320657965"
	expected := "746865206b696420646f6e277420706c6179"

	c, err := XorHexStrings(a, b)
	if err != nil {
		t.Error("For", a, "and", b, "unexpected", err)
	}

	if c != expected {
		t.Error("For", a, "and", b, "expected", expected, "got", c)
	}
}
