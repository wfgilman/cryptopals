package challenges

import (
	"testing"
)

func TestHexStringFixedXOR(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want string
	}{
		{
			a:    "1c0111001f010100061a024b53535009181c",
			b:    "686974207468652062756c6c277320657965",
			want: "746865206b696420646f6e277420706c6179",
		},
	}

	for _, test := range tests {
		want := test.want
		got := HexStringFixedXOR(test.a, test.b)
		if want != got {
			t.Fatalf("Wanted %s and got %s", want, got)
		}
	}

}
