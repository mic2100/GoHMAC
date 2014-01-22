package hmac

//Compares the HMAC in the header with the one that is generated, also checks
//to make sure that the HMAC key was generated within the expiry time
func Compare(currentHmac Message, encoder Encoder) bool {
	SetUri(currentHmac.Uri)
	SetTimestamp(currentHmac.Timestamp)
	SetAlgorithm(config.Algorithm)
	SetKey(config.Key)
	Generate(encoder)

	//does the generated HMAC differ from the current HMAC, if so return false
	if Hash.Hmac != currentHmac.Hmac {
		return false
	}

	//the generated HMAC matches the current one so return true
	return true
}
