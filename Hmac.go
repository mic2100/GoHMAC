package hmac

import (
	"fmt"
	"time"
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

type Failure struct {
	Time    time.Time
	Message string
}

func (f Failure) Error() string {
	return fmt.Sprintf("%v: %v", f.Time, f.Message)
}

//this is used so the methods can me mocked
type Encoder func(uri string, timestamp string) string

func isSafeToEncode(checkHash bool) (bool, Failure) {
	errorMsg := Failure{Time: time.Now(), Message: ""}
	errorMsg.Time = time.Now()
	passed := true
	if config.algorithm == "" {
		errorMsg.Message = "No algorithm has been set"
		passed = false
	} else if config.key == "" {
		errorMsg.Message = "No key has been set"
		passed = false
	} else if checkHash && Hash.hmac == "" {
		errorMsg.Message = "No HMAC has been set"
		passed = false
	} else if Hash.uri == "" {
		errorMsg.Message = "No URI has been set"
		passed = false
	} else if Hash.timestamp == "" {
		errorMsg.Message = "No timestamp has been set"
		passed = false
	}
	return passed, errorMsg
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
