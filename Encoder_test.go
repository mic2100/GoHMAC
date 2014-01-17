package hmac

import (
	"testing"
)

func Test_Encoder(t *testing.T) {
	algo := "sha512"
	key := "A super secret key!!"
	SetAlgorithm(algo)
	SetKey(key)

	now := "1389349028"
	uri := "/home"
	encodedString := Encode(uri, now)

	//this value will need to be recalculated if the encode method is modified
	hmacValue := "3f74e1d499ae2aea80cbbb86aa3923bfda19417e45149fb7ad474ce5143048f2dd594868d82796a9e90f2052c852850558874125b2924ffdf9a54b65f7e676ef"

	if encodedString == "" {
		t.Errorf("The HMAC was not generated (%v)", Hash.Hmac)
	}

	if encodedString != hmacValue {
		t.Errorf("The HMAC generated does not match what was expected (%v)", Hash.Hmac)
	}
}

func Benchmark_Encoder(b *testing.B) {
	algo := "sha512"
	key := "A super secret key!!"
	SetAlgorithm(algo)
	SetKey(key)

	now := "1389349028"
	uri := "/home"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Encode(uri, now)
	}
}


