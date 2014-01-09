package hmac

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

var (
	key   string
)

type configuration struct {
	algorithm string
	key       string
}

var config configuration = configuration{algorithm: "", key: ""}

type hash struct {
	hmac      string
	uri       string
	timestamp string
}

var Hash hash = hash{uri: "", timestamp: ""}

func Generate() {
	checkRequirements()
	Hash.hmac = encode(Hash.uri, Hash.timestamp)
}

func checkRequirements() {
	if config.algorithm == "" {
		panic("No algorithm has been set")
		os.Exit(1)
	} else if config.key == "" {
		panic("No key has been set")
		os.Exit(1)
	} else if Hash.uri == "" {
		panic("No URI has been set")
		os.Exit(1)
	} else if Hash.timestamp == "" {
		panic("No timestamp has been set")
		os.Exit(1)
	}
}

func encode(uri string, timestamp string) string {
	return _encode(_encode(uri + "@" + timestamp, 10) + "-" + _encode(config.key, 10), 100)
}

func _encode(encodedString string, length int) string {
	output := _enc(encodedString)
	for i := 0; i < length; i++ {
		output = _enc(output)
	}

	return output
}

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

func SetUri(uri string) {
	Hash.uri = uri
}

func SetTimestamp(ts string) {
	Hash.timestamp = ts
}

func SetAlgorithm(algo string) {
	config.algorithm = algo
}

func SetKey(key string) {
	config.key = key
}
