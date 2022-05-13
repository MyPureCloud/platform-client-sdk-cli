package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MovemanagementunitrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MovemanagementunitrequestDud struct { 
    

}

// Movemanagementunitrequest
type Movemanagementunitrequest struct { 
    // BusinessUnitId - The ID of the business unit to which to move the management unit
    BusinessUnitId string `json:"businessUnitId"`

}

// String returns a JSON representation of the model
func (o *Movemanagementunitrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Movemanagementunitrequest) MarshalJSON() ([]byte, error) {
    type Alias Movemanagementunitrequest

    if MovemanagementunitrequestMarshalled {
        return []byte("{}"), nil
    }
    MovemanagementunitrequestMarshalled = true

    return json.Marshal(&struct {
        
        BusinessUnitId string `json:"businessUnitId"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

