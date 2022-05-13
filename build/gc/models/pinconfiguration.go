package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PinconfigurationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PinconfigurationDud struct { 
    


    

}

// Pinconfiguration
type Pinconfiguration struct { 
    // MinimumLength
    MinimumLength int `json:"minimumLength"`


    // MaximumLength
    MaximumLength int `json:"maximumLength"`

}

// String returns a JSON representation of the model
func (o *Pinconfiguration) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Pinconfiguration) MarshalJSON() ([]byte, error) {
    type Alias Pinconfiguration

    if PinconfigurationMarshalled {
        return []byte("{}"), nil
    }
    PinconfigurationMarshalled = true

    return json.Marshal(&struct {
        
        MinimumLength int `json:"minimumLength"`
        
        MaximumLength int `json:"maximumLength"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

