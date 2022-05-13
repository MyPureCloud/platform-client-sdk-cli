package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContextentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContextentityDud struct { 
    

}

// Contextentity
type Contextentity struct { 
    // Name - The name of the entity.
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Contextentity) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contextentity) MarshalJSON() ([]byte, error) {
    type Alias Contextentity

    if ContextentityMarshalled {
        return []byte("{}"), nil
    }
    ContextentityMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

