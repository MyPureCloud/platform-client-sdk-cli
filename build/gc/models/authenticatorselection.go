package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuthenticatorselectionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuthenticatorselectionDud struct { 
    


    


    


    

}

// Authenticatorselection
type Authenticatorselection struct { 
    // AuthenticatorAttachment - Desired authenticator attachment modality ('platform' or 'cross-platform').
    AuthenticatorAttachment string `json:"authenticatorAttachment"`


    // RequireResidentKey - Whether a resident (discoverable) credential is required. Deprecated by the WebAuthn spec in favor of residentKey.
    RequireResidentKey bool `json:"requireResidentKey"`


    // ResidentKey - The relying party's requirement for resident (discoverable) credentials ('discouraged', 'preferred', or 'required').
    ResidentKey string `json:"residentKey"`


    // UserVerification - The user verification requirement ('discouraged', 'preferred', or 'required').
    UserVerification string `json:"userVerification"`

}

// String returns a JSON representation of the model
func (o *Authenticatorselection) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Authenticatorselection) MarshalJSON() ([]byte, error) {
    type Alias Authenticatorselection

    if AuthenticatorselectionMarshalled {
        return []byte("{}"), nil
    }
    AuthenticatorselectionMarshalled = true

    return json.Marshal(&struct {
        
        AuthenticatorAttachment string `json:"authenticatorAttachment"`
        
        RequireResidentKey bool `json:"requireResidentKey"`
        
        ResidentKey string `json:"residentKey"`
        
        UserVerification string `json:"userVerification"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

