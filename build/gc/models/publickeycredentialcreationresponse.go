package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PublickeycredentialcreationresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PublickeycredentialcreationresponseDud struct { 
    


    


    


    


    


    

}

// Publickeycredentialcreationresponse
type Publickeycredentialcreationresponse struct { 
    // Id - The credential identifier (base64url-encoded).
    Id string `json:"id"`


    // VarType - The credential type (must be 'public-key').
    VarType string `json:"type"`


    // RawId - The raw credential identifier as a binary value (base64url-encoded).
    RawId string `json:"rawId"`


    // AuthenticatorAttachment - The authenticator attachment modality used ('platform' or 'cross-platform').
    AuthenticatorAttachment string `json:"authenticatorAttachment"`


    // ClientExtensionResults - Outputs from client-side WebAuthn extensions.
    ClientExtensionResults map[string]interface{} `json:"clientExtensionResults"`


    // Response - The authenticator's attestation response.
    Response Authenticatorattestationresponse `json:"response"`

}

// String returns a JSON representation of the model
func (o *Publickeycredentialcreationresponse) String() string {
    
    
    
    
     o.ClientExtensionResults = map[string]interface{}{"": Interface{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Publickeycredentialcreationresponse) MarshalJSON() ([]byte, error) {
    type Alias Publickeycredentialcreationresponse

    if PublickeycredentialcreationresponseMarshalled {
        return []byte("{}"), nil
    }
    PublickeycredentialcreationresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        VarType string `json:"type"`
        
        RawId string `json:"rawId"`
        
        AuthenticatorAttachment string `json:"authenticatorAttachment"`
        
        ClientExtensionResults map[string]interface{} `json:"clientExtensionResults"`
        
        Response Authenticatorattestationresponse `json:"response"`
        *Alias
    }{

        


        


        


        


        
        ClientExtensionResults: map[string]interface{}{"": Interface{}},
        


        

        Alias: (*Alias)(u),
    })
}

