package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SharedentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SharedentityDud struct { 
    

}

// Sharedentity
type Sharedentity struct { 
    // Id
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Sharedentity) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sharedentity) MarshalJSON() ([]byte, error) {
    type Alias Sharedentity

    if SharedentityMarshalled {
        return []byte("{}"), nil
    }
    SharedentityMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

