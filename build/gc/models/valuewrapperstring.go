package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ValuewrapperstringMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ValuewrapperstringDud struct { 
    

}

// Valuewrapperstring - An object to provide context to nullable fields in PATCH requests
type Valuewrapperstring struct { 
    // Value - The value for the associated field
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Valuewrapperstring) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Valuewrapperstring) MarshalJSON() ([]byte, error) {
    type Alias Valuewrapperstring

    if ValuewrapperstringMarshalled {
        return []byte("{}"), nil
    }
    ValuewrapperstringMarshalled = true

    return json.Marshal(&struct { 
        Value string `json:"value"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

