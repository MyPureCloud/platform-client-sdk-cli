package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PolicyupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PolicyupdateDud struct { 
    

}

// Policyupdate
type Policyupdate struct { 
    // Enabled
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Policyupdate) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Policyupdate) MarshalJSON() ([]byte, error) {
    type Alias Policyupdate

    if PolicyupdateMarshalled {
        return []byte("{}"), nil
    }
    PolicyupdateMarshalled = true

    return json.Marshal(&struct { 
        Enabled bool `json:"enabled"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

