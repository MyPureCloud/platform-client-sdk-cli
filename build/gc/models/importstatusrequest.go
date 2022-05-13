package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ImportstatusrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ImportstatusrequestDud struct { 
    

}

// Importstatusrequest
type Importstatusrequest struct { 
    // Status - New status for existing import operation
    Status string `json:"status"`

}

// String returns a JSON representation of the model
func (o *Importstatusrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Importstatusrequest) MarshalJSON() ([]byte, error) {
    type Alias Importstatusrequest

    if ImportstatusrequestMarshalled {
        return []byte("{}"), nil
    }
    ImportstatusrequestMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

