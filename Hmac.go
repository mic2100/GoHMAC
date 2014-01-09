package hmac

import (
	"os"
)

var key string

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

var Hash = hash{hmac: "", uri: "", timestamp: ""}

func checkRequirements(checkHash bool) {
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

	if checkHash && Hash.hmac == "" {
		panic("No HMAC has been set")
		os.Exit(1)
	}
}

func SetHmac(hmac string) {
	Hash.hmac = hmac
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
