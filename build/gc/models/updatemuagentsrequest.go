package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatemuagentsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatemuagentsrequestDud struct { 
    

}

// Updatemuagentsrequest
type Updatemuagentsrequest struct { 
    // Entities - List of agents to update
    Entities []Updatemuagentrequest `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Updatemuagentsrequest) String() string {
     o.Entities = []Updatemuagentrequest{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatemuagentsrequest) MarshalJSON() ([]byte, error) {
    type Alias Updatemuagentsrequest

    if UpdatemuagentsrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdatemuagentsrequestMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Updatemuagentrequest `json:"entities"`
        *Alias
    }{

        
        Entities: []Updatemuagentrequest{{}},
        

        Alias: (*Alias)(u),
    })
}

