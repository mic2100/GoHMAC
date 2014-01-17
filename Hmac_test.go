package hmac

import (
	"testing"
)

func Test_IsSafeToEncode_NoAlgorithm(t *testing.T) {
	Hash = hash{Hmac: "", Uri: "", Timestamp: ""}
	config = configuration{Algorithm: "", Key: ""}
	passed, failure := isSafeToEncode(false)

	if passed {
		t.Errorf("Failed to check if the algorithm has been set, the boolean was an incorrect value")
	}

	if failure.Message != "No algorithm has been set" {
		t.Errorf("Failed to check if the algorithm has been set, incorrect error message")
	}
}

func Test_IsSafeToEncode_NoKey(t *testing.T) {
	Hash = hash{Hmac: "", Uri: "", Timestamp: ""}
	config = configuration{Algorithm: "sha512", Key: ""}
	passed, failure := isSafeToEncode(false)

	if passed {
		t.Errorf("Failed to check if the key has been set, the boolean was an incorrect value")
	}

	if failure.Message != "No key has been set" {
		t.Errorf("Failed to check if the key has been set, incorrect error message")
	}
}

func Test_IsSafeToEncode_NoHmac(t *testing.T) {
	Hash = hash{Hmac: "", Uri: "", Timestamp: ""}
	config = configuration{Algorithm: "sha512", Key: "12345"}
	passed, failure := isSafeToEncode(true)
	if passed {
		t.Errorf("Failed to check if the HMAC has been set, the boolean was an incorrect value")
	}

	if failure.Message != "No HMAC has been set" {
		t.Errorf("Failed to check if the HMAC has been set, incorrect error message")
	}
}

func Test_IsSafeToEncode_NoUri(t *testing.T) {
	Hash = hash{Hmac: "54321", Uri: "", Timestamp: ""}
	config = configuration{Algorithm: "sha512", Key: "12345"}
	passed, failure := isSafeToEncode(false)

	if passed {
		t.Errorf("Failed to check if the URI has been set, the boolean was an incorrect value")
	}

	if failure.Message != "No URI has been set" {
		t.Errorf("Failed to check if the URI has been set, incorrect error message")
	}
}

func Test_IsSafeToEncode_NoTimestamp(t *testing.T) {
	Hash = hash{Hmac: "54321", Uri: "/home", Timestamp: ""}
	config = configuration{Algorithm: "sha512", Key: "12345"}
	passed, failure := isSafeToEncode(false)

	if passed {
		t.Errorf("Failed to check if the timestamp has been set, the boolean was an incorrect value")
	}

	if failure.Message != "No timestamp has been set" {
		t.Errorf("Failed to check if the timestamp has been set, incorrect error message")
	}
}

func Test_IsSafeToEncode_Passed(t *testing.T) {
	Hash = hash{Hmac: "54321", Uri: "/home", Timestamp: "1234567890"}
	config = configuration{Algorithm: "sha512", Key: "12345"}
	passed, failure := isSafeToEncode(false)

	if !passed {
		t.Errorf("Failed to check if the IsSafeToEncode func returns true")
	}

	if failure.Message != "" {
		t.Errorf("Failed to check if the IsSafeToEncode has the correct error message")
	}
}
