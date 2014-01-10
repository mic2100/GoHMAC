package hmac

//Compares the HMAC in the header with the one that is generated, also checks
//to make sure that the HMAC key was generated within the expiry time
func Compare(currentHmac hash) bool {
	SetUri(currentHmac.uri)
	SetTimestamp(currentHmac.timestamp)
	SetAlgorithm(config.algorithm)
	SetKey(config.key)
	Generate()

	//does the generated HMAC differ from the current HMAC, if so return false
	if Hash.hmac != currentHmac.hmac {
		return false
	}

	//the generated HMAC matches the current one so return true
	return true
}
