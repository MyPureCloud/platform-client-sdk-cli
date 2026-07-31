package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CredentialparameterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CredentialparameterDud struct { 
    


    

}

// Credentialparameter
type Credentialparameter struct { 
    // VarType - The public key credential type (e.g., 'public-key').
    VarType string `json:"type"`


    // Alg - The COSE algorithm identifier (e.g., -7 for ES256, -257 for RS256).
    Alg int `json:"alg"`

}

// String returns a JSON representation of the model
func (o *Credentialparameter) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Credentialparameter) MarshalJSON() ([]byte, error) {
    type Alias Credentialparameter

    if CredentialparameterMarshalled {
        return []byte("{}"), nil
    }
    CredentialparameterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Alg int `json:"alg"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

