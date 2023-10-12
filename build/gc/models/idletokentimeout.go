package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IdletokentimeoutMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IdletokentimeoutDud struct { 
    


    

}

// Idletokentimeout
type Idletokentimeout struct { 
    // IdleTokenTimeoutSeconds - Token timeout length in seconds. Must be at least 5 minutes and 8 hours or less (if HIPAA is disabled) or 15 minutes or less (if HIPAA is enabled).
    IdleTokenTimeoutSeconds int `json:"idleTokenTimeoutSeconds"`


    // EnableIdleTokenTimeout - Indicates whether the Token Timeout should be enabled or disabled.
    EnableIdleTokenTimeout bool `json:"enableIdleTokenTimeout"`

}

// String returns a JSON representation of the model
func (o *Idletokentimeout) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Idletokentimeout) MarshalJSON() ([]byte, error) {
    type Alias Idletokentimeout

    if IdletokentimeoutMarshalled {
        return []byte("{}"), nil
    }
    IdletokentimeoutMarshalled = true

    return json.Marshal(&struct {
        
        IdleTokenTimeoutSeconds int `json:"idleTokenTimeoutSeconds"`
        
        EnableIdleTokenTimeout bool `json:"enableIdleTokenTimeout"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

