package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactenrichoptionsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactenrichoptionsDud struct { 
    

}

// Contactenrichoptions
type Contactenrichoptions struct { 
    // Promote - If true, any matched contact will have its type updated to Curated. If false, any matched contact will keep its original type in the operation.
    Promote bool `json:"promote"`

}

// String returns a JSON representation of the model
func (o *Contactenrichoptions) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactenrichoptions) MarshalJSON() ([]byte, error) {
    type Alias Contactenrichoptions

    if ContactenrichoptionsMarshalled {
        return []byte("{}"), nil
    }
    ContactenrichoptionsMarshalled = true

    return json.Marshal(&struct {
        
        Promote bool `json:"promote"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

