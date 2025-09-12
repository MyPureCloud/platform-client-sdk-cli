package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RemoveentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RemoveentityDud struct { 
    

}

// Removeentity
type Removeentity struct { 
    // Id - Unique identifier of the entity to be removed
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Removeentity) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Removeentity) MarshalJSON() ([]byte, error) {
    type Alias Removeentity

    if RemoveentityMarshalled {
        return []byte("{}"), nil
    }
    RemoveentityMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

