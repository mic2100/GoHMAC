package hmac

import "os"

//Generates the HMAC and puts all the generated values in the Hash struct
func Generate(encoder Encoder) {
	passed, failure := isSafeToEncode(false)
	if !passed {
		panic(failure.Error())
		os.Exit(1)
	}
	Hash.Hmac = encoder(Hash.Uri, Hash.Timestamp)
}
