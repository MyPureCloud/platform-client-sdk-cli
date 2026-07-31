package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BeginwebauthnregistrationresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BeginwebauthnregistrationresponseDud struct { 
    

}

// Beginwebauthnregistrationresponse
type Beginwebauthnregistrationresponse struct { 
    // PublicKey - The public key credential creation options the client should pass to the WebAuthn API's navigator.credentials.create() call.
    PublicKey Publickeycredentialcreationoptions `json:"publicKey"`

}

// String returns a JSON representation of the model
func (o *Beginwebauthnregistrationresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Beginwebauthnregistrationresponse) MarshalJSON() ([]byte, error) {
    type Alias Beginwebauthnregistrationresponse

    if BeginwebauthnregistrationresponseMarshalled {
        return []byte("{}"), nil
    }
    BeginwebauthnregistrationresponseMarshalled = true

    return json.Marshal(&struct {
        
        PublicKey Publickeycredentialcreationoptions `json:"publicKey"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

