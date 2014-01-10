package hmac

import (
	"testing"
)

func Test_Generate(t *testing.T) {
	now := "1389349028"
	uri := "/home"
	algo := "sha512"
	key := "A super secret key!!"
	SetUri(uri)
	SetTimestamp(now)
	SetAlgorithm(algo)
	SetKey(key)
	Generate(encoder)

	//this value will need to be recalculated if the encode method is modified
	hmacValue := "3f74e1d499ae2aea80cbbb86aa3923bfda19417e45149fb7ad474ce5143048f2dd594868d82796a9e90f2052c852850558874125b2924ffdf9a54b65f7e676ef"

	if Hash.hmac == "" {
		t.Errorf("The HMAC was not generated (%v)", Hash.hmac)
	}

	if Hash.hmac != hmacValue {
		t.Errorf("The HMAC generated does not match what was expected (%v)", Hash.hmac)
	}

	if Hash.timestamp != now {
		t.Errorf("The timestamp appears to have been changed, expecting: %s received: (%v)", now, Hash.timestamp)
	}

	if Hash.uri != uri {
		t.Errorf("The URI appears to have been changed, expecting: %s received: (%v)", uri, Hash.uri)
	}

}

func Benchmark_Generate(b *testing.B) {
	now := "1389349028"
	uri := "/home"
	algo := "sha512"
	key := "A super secret key!!"
	SetUri(uri)
	SetTimestamp(now)
	SetAlgorithm(algo)
	SetKey(key)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Generate(mock_encoder)
	}
}

func mock_encoder(uri string, timestamp string) string {
	return "abcdefghijklmnopqrstuvwxyz1234567890"
}
