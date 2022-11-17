package challenges

import "testing"

func TestHexToBase64(t *testing.T) {
	hexValue := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	want := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	got := HexToBase64(hexValue)
	if want != got {
		t.Fatalf("Wanted %s and got %s", want, got)
	}
}
