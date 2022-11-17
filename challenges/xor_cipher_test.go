package challenges

import (
	"math"
	"testing"
)

func TestXORCipher(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{
			s:    "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736",
			want: "X",
		},
	}

	for _, test := range tests {
		_, got := XORCipher(test.s)
		if test.want != got {
			t.Fatalf("Wanted %s and got %s", test.want, got)
		}
	}
}

func TestScoreText(t *testing.T) {
	const training = "You're simply the best, better than all the rest, better than anyone, better than all the rest"

	tests := []struct {
		text string
		want float64
	}{
		{
			text: "I'm stuck on your heart",
			want: float64(15.383441258094356),
		},
		{
			text: "Paʻa wau i kou puʻuwai",
			want: float64(10.280141843971633),
		},
		{
			text: "Sunt blocat în inima ta",
			want: float64(13.49468085106383),
		},
	}

	for _, test := range tests {
		got := ScoreText(training, test.text)
		if math.Abs(test.want-got) > 1e-6 {
			t.Fatalf("Wanted %f and got %f", test.want, got)
		}
	}
}
