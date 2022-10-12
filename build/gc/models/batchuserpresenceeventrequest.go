package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BatchuserpresenceeventrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BatchuserpresenceeventrequestDud struct { 
    

}

// Batchuserpresenceeventrequest - A maximum of 100 events are allowed per request
type Batchuserpresenceeventrequest struct { 
    // UserPresenceEvents - UserPresence events for this batch
    UserPresenceEvents []Userpresenceevent `json:"userPresenceEvents"`

}

// String returns a JSON representation of the model
func (o *Batchuserpresenceeventrequest) String() string {
     o.UserPresenceEvents = []Userpresenceevent{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Batchuserpresenceeventrequest) MarshalJSON() ([]byte, error) {
    type Alias Batchuserpresenceeventrequest

    if BatchuserpresenceeventrequestMarshalled {
        return []byte("{}"), nil
    }
    BatchuserpresenceeventrequestMarshalled = true

    return json.Marshal(&struct {
        
        UserPresenceEvents []Userpresenceevent `json:"userPresenceEvents"`
        *Alias
    }{

        
        UserPresenceEvents: []Userpresenceevent{{}},
        

        Alias: (*Alias)(u),
    })
}

