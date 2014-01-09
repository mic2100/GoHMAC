package hmac

//Generates the HMAC and puts all the generated values in the Hash struct
func Generate() {
	checkRequirements(false)
	Hash.hmac = encode(Hash.uri, Hash.timestamp)
}
