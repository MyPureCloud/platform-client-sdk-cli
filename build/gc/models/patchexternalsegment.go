package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchexternalsegmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchexternalsegmentDud struct { 
    

}

// Patchexternalsegment
type Patchexternalsegment struct { 
    // Name - Name for the external segment in the system where it originates from.
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Patchexternalsegment) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Patchexternalsegment) MarshalJSON() ([]byte, error) {
    type Alias Patchexternalsegment

    if PatchexternalsegmentMarshalled {
        return []byte("{}"), nil
    }
    PatchexternalsegmentMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

