package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DataschemarefMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DataschemarefDud struct { 
    

}

// Dataschemaref
type Dataschemaref struct { 
    // Ref
    Ref string `json:"$ref"`

}

// String returns a JSON representation of the model
func (o *Dataschemaref) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dataschemaref) MarshalJSON() ([]byte, error) {
    type Alias Dataschemaref

    if DataschemarefMarshalled {
        return []byte("{}"), nil
    }
    DataschemarefMarshalled = true

    return json.Marshal(&struct {
        
        Ref string `json:"$ref"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

