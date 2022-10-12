package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BatchuserroutingstatuseventrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BatchuserroutingstatuseventrequestDud struct { 
    

}

// Batchuserroutingstatuseventrequest - A maximum of 100 events are allowed per request
type Batchuserroutingstatuseventrequest struct { 
    // UserRoutingStatusEvents - UserRoutingStatus events for this batch
    UserRoutingStatusEvents []Userroutingstatusevent `json:"userRoutingStatusEvents"`

}

// String returns a JSON representation of the model
func (o *Batchuserroutingstatuseventrequest) String() string {
     o.UserRoutingStatusEvents = []Userroutingstatusevent{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Batchuserroutingstatuseventrequest) MarshalJSON() ([]byte, error) {
    type Alias Batchuserroutingstatuseventrequest

    if BatchuserroutingstatuseventrequestMarshalled {
        return []byte("{}"), nil
    }
    BatchuserroutingstatuseventrequestMarshalled = true

    return json.Marshal(&struct {
        
        UserRoutingStatusEvents []Userroutingstatusevent `json:"userRoutingStatusEvents"`
        *Alias
    }{

        
        UserRoutingStatusEvents: []Userroutingstatusevent{{}},
        

        Alias: (*Alias)(u),
    })
}

