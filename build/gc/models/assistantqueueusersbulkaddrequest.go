package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AssistantqueueusersbulkaddrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AssistantqueueusersbulkaddrequestDud struct { 
    

}

// Assistantqueueusersbulkaddrequest
type Assistantqueueusersbulkaddrequest struct { 
    // Entities - List of users to assign assistant for. Maximum users to add is 100 per request.
    Entities []Assistantqueueuser `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Assistantqueueusersbulkaddrequest) String() string {
     o.Entities = []Assistantqueueuser{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Assistantqueueusersbulkaddrequest) MarshalJSON() ([]byte, error) {
    type Alias Assistantqueueusersbulkaddrequest

    if AssistantqueueusersbulkaddrequestMarshalled {
        return []byte("{}"), nil
    }
    AssistantqueueusersbulkaddrequestMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Assistantqueueuser `json:"entities"`
        *Alias
    }{

        
        Entities: []Assistantqueueuser{{}},
        

        Alias: (*Alias)(u),
    })
}

