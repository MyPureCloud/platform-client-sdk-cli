package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuthenticatorattestationresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuthenticatorattestationresponseDud struct { 
    


    


    

}

// Authenticatorattestationresponse
type Authenticatorattestationresponse struct { 
    // ClientDataJSON - The JSON-serialized client data passed to the authenticator (base64url-encoded).
    ClientDataJSON string `json:"clientDataJSON"`


    // AttestationObject - The attestation object containing the credential public key and attestation statement (base64url-encoded CBOR).
    AttestationObject string `json:"attestationObject"`


    // Transports - The transports the authenticator is believed to support.
    Transports []string `json:"transports"`

}

// String returns a JSON representation of the model
func (o *Authenticatorattestationresponse) String() string {
    
    
     o.Transports = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Authenticatorattestationresponse) MarshalJSON() ([]byte, error) {
    type Alias Authenticatorattestationresponse

    if AuthenticatorattestationresponseMarshalled {
        return []byte("{}"), nil
    }
    AuthenticatorattestationresponseMarshalled = true

    return json.Marshal(&struct {
        
        ClientDataJSON string `json:"clientDataJSON"`
        
        AttestationObject string `json:"attestationObject"`
        
        Transports []string `json:"transports"`
        *Alias
    }{

        


        


        
        Transports: []string{""},
        

        Alias: (*Alias)(u),
    })
}

