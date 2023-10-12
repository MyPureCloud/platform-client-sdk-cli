package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PinnedmessagerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PinnedmessagerequestDud struct { 
    

}

// Pinnedmessagerequest
type Pinnedmessagerequest struct { 
    // PinnedMessageIds - Ids of the messages to pin
    PinnedMessageIds []string `json:"pinnedMessageIds"`

}

// String returns a JSON representation of the model
func (o *Pinnedmessagerequest) String() string {
     o.PinnedMessageIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Pinnedmessagerequest) MarshalJSON() ([]byte, error) {
    type Alias Pinnedmessagerequest

    if PinnedmessagerequestMarshalled {
        return []byte("{}"), nil
    }
    PinnedmessagerequestMarshalled = true

    return json.Marshal(&struct {
        
        PinnedMessageIds []string `json:"pinnedMessageIds"`
        *Alias
    }{

        
        PinnedMessageIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

