package hmac

import (
	"testing"
)

func Test_CompareFalse(t *testing.T) {
	now := "1389349028"
	uri := "/home"
	algo := "sha512"
	key := "A super secret key!!"
	SetUri(uri)
	SetTimestamp(now)
	SetAlgorithm(algo)
	SetKey(key)
	Generate(encoder)

	SetHmac("12345")

	result := Compare(Hash, encoder)

	if result {
		t.Errorf("Campare is true when it should be false (%+v)", Hash)
	}

}

func Test_ComapreTrue(t *testing.T) {
	now := "1389349028"
	uri := "/home"
	algo := "sha512"
	key := "A super secret key!!"
	SetUri(uri)
	SetTimestamp(now)
	SetAlgorithm(algo)
	SetKey(key)
	Generate(encoder)

	result := Compare(Hash, encoder)

	if !result {
		t.Errorf("Campare is false when it should be true (%+v)", Hash)
	}
}

func Benchmark_Compare(b *testing.B) {
	now := "1389349028"
	uri := "/home"
	algo := "sha512"
	key := "A super secret key!!"
	SetUri(uri)
	SetTimestamp(now)
	SetAlgorithm(algo)
	SetKey(key)
	Generate(mock_encoder)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Compare(Hash, mock_encoder)
	}
}
