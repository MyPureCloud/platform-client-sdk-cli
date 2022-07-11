package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CopyworkplanMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CopyworkplanDud struct { 
    

}

// Copyworkplan
type Copyworkplan struct { 
    // Name - Name of the copied work plan
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Copyworkplan) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Copyworkplan) MarshalJSON() ([]byte, error) {
    type Alias Copyworkplan

    if CopyworkplanMarshalled {
        return []byte("{}"), nil
    }
    CopyworkplanMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

