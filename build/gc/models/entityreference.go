package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EntityreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EntityreferenceDud struct { 
    

}

// Entityreference
type Entityreference struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Entityreference) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Entityreference) MarshalJSON() ([]byte, error) {
    type Alias Entityreference

    if EntityreferenceMarshalled {
        return []byte("{}"), nil
    }
    EntityreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

