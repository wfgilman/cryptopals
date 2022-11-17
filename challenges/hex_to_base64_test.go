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

func TestEncodeHexToString(t *testing.T) {
	want := "746865206b696420646f6e277420706c6179"
	got := EncodeHexToString([]byte("the kid don't play"))
	if want != got {
		t.Fatalf("Wanted %s and got %s", want, got)
	}
}

func TestDecodeHexToASCIIString(t *testing.T) {
	tests := []struct {
		hex   string
		ascii string
	}{
		{
			hex:   "41696E2774206E6F207061727479206C696B652061207765737420636F617374207061727479",
			ascii: "Ain't no party like a west coast party",
		},
		{
			hex:   "4C6F646920646F64692C207765206C696B6520746F207061727479",
			ascii: "Lodi dodi, we like to party",
		},
	}

	for _, test := range tests {
		want := test.ascii
		got := DecodeHexString(test.hex)
		if want != string(got) {
			t.Fatalf("Wanted %s and got %s", want, got)
		}
	}
}
