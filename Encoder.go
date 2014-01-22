package hmac

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
)

//Encodes the HMAC based on the URI, Timestamp and Key
func Encode(uri string, timestamp string) string {
	return _encode(_encode(uri+"@"+timestamp, 10)+"-"+_encode(config.Key, 10), 100)
}

//Encodes the value for the number of iterations sent in the method call
func _encode(encodedString string, numOfIterations int) string {
	output := _enc(encodedString)
	for i := 0; i < numOfIterations; i++ {
		output = _enc(output)
	}

	return output
}

//Encodes the value based on the selected algorithm
func _enc(value string) string {
	var encodedString string
	switch config.Algorithm {
	case "sha512":
		encodedString = _hash(sha512.New(), value)
	case "sha256":
		encodedString = _hash(sha256.New(), value)
	}

	return encodedString
}

func _hash(hasher hash.Hash, value string) string {
	hasher.Write([]byte(value))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}
