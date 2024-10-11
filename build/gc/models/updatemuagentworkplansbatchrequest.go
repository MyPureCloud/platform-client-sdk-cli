package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatemuagentworkplansbatchrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatemuagentworkplansbatchrequestDud struct { 
    

}

// Updatemuagentworkplansbatchrequest
type Updatemuagentworkplansbatchrequest struct { 
    // Entities
    Entities []Updatemuagentworkplanrequest `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Updatemuagentworkplansbatchrequest) String() string {
     o.Entities = []Updatemuagentworkplanrequest{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatemuagentworkplansbatchrequest) MarshalJSON() ([]byte, error) {
    type Alias Updatemuagentworkplansbatchrequest

    if UpdatemuagentworkplansbatchrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdatemuagentworkplansbatchrequestMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Updatemuagentworkplanrequest `json:"entities"`
        *Alias
    }{

        
        Entities: []Updatemuagentworkplanrequest{{}},
        

        Alias: (*Alias)(u),
    })
}

