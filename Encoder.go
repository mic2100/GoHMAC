package hmac

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

//Encodes the HMAC based on the URI, Timestamp and Key
func encoder(uri string, timestamp string) string {
	return _encode(_encode(uri+"@"+timestamp, 10)+"-"+_encode(config.key, 10), 100)
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
	switch config.algorithm {
	case "sha512":
		hasher := sha512.New()
		hasher.Write([]byte(value))
		encodedString = fmt.Sprintf("%x", hasher.Sum(nil))
	case "sha256":
		hasher := sha256.New()
		hasher.Write([]byte(value))
		encodedString = fmt.Sprintf("%x", hasher.Sum(nil))
	}

	return encodedString
}
