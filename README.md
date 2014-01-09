# Instructions

## Installing

To install and use this package you need to run

    $ go get github.com/mic2100/GoHMAC
    
from the `src` directory of your project. Then in the `import ()` include the library

    package sample

    import (
        "github.com/mic2100/GoHMAC"
    )
   
 
## Usage
 
### Generation

To generate an HMAC:

    hmac.SetUri("/home")
    hmac.SetTimestamp(strconv.FormatInt(time.Now().Unix(), 10))
    hmac.SetAlgorithm("sha512") //can be "sha512" or "sha256" more will be added later
    hmac.SetKey("A super secret key!!!") //must be kept secure because this will make everything work
    hmac.Generate()
    generatedHmac := hmac.Hash
    
    fmt.Printf("%+v", generatedHmac)
    
Will output the unique generated HMAC to the command line

    $ {hmac:1e2c240264cf8fd84f442f35a1de42895ecd5217143313787c1261eeceb8bc5f7972ac8e6539a96a466e7dd46d0753b63a9dd402465b4f04b9184b55ab288eb9 uri:/home timestamp:1389267365}
    
This information can then be sent in the header of the request to be used to authenticate the request.

### Compare

To compare an HMAC:

    hmac.SetHmac(hmac) //the key extracted from the header of the request
    hmac.SetUri(uri) //the uri from the header
    hmac.SetTimestamp(timestamp) //the timestamp used to create this hash
    hmac.SetAlgorithm("sha512")
    hmac.SetKey("A super secret key!!!")
    
    if hmac.Compare(hmac.Hash) {
        //this hash was invalid
    }
    
    //this hash was valid