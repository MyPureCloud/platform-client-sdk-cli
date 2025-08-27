package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AssistantqueueuserMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AssistantqueueuserDud struct { 
    

}

// Assistantqueueuser
type Assistantqueueuser struct { 
    // Id - The globally unique identifier for the user.
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Assistantqueueuser) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Assistantqueueuser) MarshalJSON() ([]byte, error) {
    type Alias Assistantqueueuser

    if AssistantqueueuserMarshalled {
        return []byte("{}"), nil
    }
    AssistantqueueuserMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

