package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OneononeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OneononeDud struct { 
    

}

// Oneonone
type Oneonone struct { 
    // PinnedMessages - Room's pinned messages
    PinnedMessages []Addressableentityref `json:"pinnedMessages"`

}

// String returns a JSON representation of the model
func (o *Oneonone) String() string {
     o.PinnedMessages = []Addressableentityref{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Oneonone) MarshalJSON() ([]byte, error) {
    type Alias Oneonone

    if OneononeMarshalled {
        return []byte("{}"), nil
    }
    OneononeMarshalled = true

    return json.Marshal(&struct {
        
        PinnedMessages []Addressableentityref `json:"pinnedMessages"`
        *Alias
    }{

        
        PinnedMessages: []Addressableentityref{{}},
        

        Alias: (*Alias)(u),
    })
}

