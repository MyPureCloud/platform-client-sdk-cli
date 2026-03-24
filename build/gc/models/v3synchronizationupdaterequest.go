package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3synchronizationupdaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3synchronizationupdaterequestDud struct { 
    

}

// V3synchronizationupdaterequest
type V3synchronizationupdaterequest struct { 
    // Status - The status of the synchronization.
    Status string `json:"status"`

}

// String returns a JSON representation of the model
func (o *V3synchronizationupdaterequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3synchronizationupdaterequest) MarshalJSON() ([]byte, error) {
    type Alias V3synchronizationupdaterequest

    if V3synchronizationupdaterequestMarshalled {
        return []byte("{}"), nil
    }
    V3synchronizationupdaterequestMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

