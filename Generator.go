package hmac

//Generates the HMAC and puts all the generated values in the Hash struct
func Generate(encoder Encoder) {
	checkRequirements(false)
	Hash.hmac = encoder(Hash.uri, Hash.timestamp)
}
