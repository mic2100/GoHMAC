package hmac

import (
	"fmt"
	"time"
)

var key string

type configuration struct {
	Algorithm string
	Key       string
}

var config configuration = configuration{Algorithm: "", Key: ""}

type hash struct {
	Hmac      string
	Uri       string
	Timestamp string
}

var Hash = hash{Hmac: "", Uri: "", Timestamp: ""}

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
	if config.Algorithm == "" {
		errorMsg.Message = "No algorithm has been set"
		passed = false
	} else if config.Key == "" {
		errorMsg.Message = "No key has been set"
		passed = false
	} else if checkHash && Hash.Hmac == "" {
		errorMsg.Message = "No HMAC has been set"
		passed = false
	} else if Hash.Uri == "" {
		errorMsg.Message = "No URI has been set"
		passed = false
	} else if Hash.Timestamp == "" {
		errorMsg.Message = "No timestamp has been set"
		passed = false
	}
	return passed, errorMsg
}

func SetHmac(hmac string) {
	Hash.Hmac = hmac
}

func SetUri(uri string) {
	Hash.Uri = uri
}

func SetTimestamp(ts string) {
	Hash.Timestamp = ts
}

func SetAlgorithm(algo string) {
	config.Algorithm = algo
}

func SetKey(key string) {
	config.Key = key
}
