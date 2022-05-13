package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchbureschedulingoptionsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchbureschedulingoptionsrequestDud struct { 
    

}

// Patchbureschedulingoptionsrequest
type Patchbureschedulingoptionsrequest struct { 
    // ManagementUnits - Per-management unit rescheduling options to update
    ManagementUnits []Patchbureschedulingoptionsmanagementunitrequest `json:"managementUnits"`

}

// String returns a JSON representation of the model
func (o *Patchbureschedulingoptionsrequest) String() string {
     o.ManagementUnits = []Patchbureschedulingoptionsmanagementunitrequest{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Patchbureschedulingoptionsrequest) MarshalJSON() ([]byte, error) {
    type Alias Patchbureschedulingoptionsrequest

    if PatchbureschedulingoptionsrequestMarshalled {
        return []byte("{}"), nil
    }
    PatchbureschedulingoptionsrequestMarshalled = true

    return json.Marshal(&struct {
        
        ManagementUnits []Patchbureschedulingoptionsmanagementunitrequest `json:"managementUnits"`
        *Alias
    }{

        
        ManagementUnits: []Patchbureschedulingoptionsmanagementunitrequest{{}},
        

        Alias: (*Alias)(u),
    })
}

