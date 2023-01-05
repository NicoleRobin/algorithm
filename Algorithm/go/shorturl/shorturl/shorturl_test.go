package shorturl

import "testing"

func TestShortUrl(t *testing.T) {
	tests := []struct {
		id       int
		expected string
	}{
		{
			id:       0,
			expected: "000000",
		},
	}

	for _, tt := range tests {
		res := ShortUrl(tt.id)
		if res != tt.expected {
			t.Fatalf("res:%s not equal to expectd:%s", res, tt.expected)
		}
	}
}

func BenchmarkShortUrl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShortUrl(i)
	}
}
