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
		outputError("No algorithm has been set")
	} else if config.key == "" {
		outputError("No key has been set")
	} else if checkHash && Hash.hmac == "" {
		outputError("No HMAC has been set")
	} else if Hash.uri == "" {
		outputError("No URI has been set")
	} else if Hash.timestamp == "" {
		outputError("No timestamp has been set")
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

func outputError(message string) {
	panic(message)
	os.Exit(1)
}
