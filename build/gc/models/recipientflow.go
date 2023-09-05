package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecipientflowMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecipientflowDud struct { 
    

}

// Recipientflow
type Recipientflow struct { 
    // Id - The flow identifier
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Recipientflow) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recipientflow) MarshalJSON() ([]byte, error) {
    type Alias Recipientflow

    if RecipientflowMarshalled {
        return []byte("{}"), nil
    }
    RecipientflowMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

