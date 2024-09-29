package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SyncstatusrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SyncstatusrequestDud struct { 
    

}

// Syncstatusrequest
type Syncstatusrequest struct { 
    // Status - New status for an existing sync operation
    Status string `json:"status"`

}

// String returns a JSON representation of the model
func (o *Syncstatusrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Syncstatusrequest) MarshalJSON() ([]byte, error) {
    type Alias Syncstatusrequest

    if SyncstatusrequestMarshalled {
        return []byte("{}"), nil
    }
    SyncstatusrequestMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

