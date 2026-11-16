package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CredentialpropertiesoutputMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CredentialpropertiesoutputDud struct { 
    

}

// Credentialpropertiesoutput
type Credentialpropertiesoutput struct { 
    // Rk - Whether the credential is a resident (discoverable) credential.
    Rk bool `json:"rk"`

}

// String returns a JSON representation of the model
func (o *Credentialpropertiesoutput) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Credentialpropertiesoutput) MarshalJSON() ([]byte, error) {
    type Alias Credentialpropertiesoutput

    if CredentialpropertiesoutputMarshalled {
        return []byte("{}"), nil
    }
    CredentialpropertiesoutputMarshalled = true

    return json.Marshal(&struct {
        
        Rk bool `json:"rk"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

