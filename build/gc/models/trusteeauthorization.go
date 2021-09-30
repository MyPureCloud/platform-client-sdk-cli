package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TrusteeauthorizationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TrusteeauthorizationDud struct { 
    Permissions []string `json:"permissions"`

}

// Trusteeauthorization
type Trusteeauthorization struct { 
    

}

// String returns a JSON representation of the model
func (o *Trusteeauthorization) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trusteeauthorization) MarshalJSON() ([]byte, error) {
    type Alias Trusteeauthorization

    if TrusteeauthorizationMarshalled {
        return []byte("{}"), nil
    }
    TrusteeauthorizationMarshalled = true

    return json.Marshal(&struct { 
        
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

