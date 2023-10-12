package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemmanualassignMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemmanualassignDud struct { 
    

}

// Workitemmanualassign
type Workitemmanualassign struct { 
    // Id - The globally unique identifier for this user.
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Workitemmanualassign) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemmanualassign) MarshalJSON() ([]byte, error) {
    type Alias Workitemmanualassign

    if WorkitemmanualassignMarshalled {
        return []byte("{}"), nil
    }
    WorkitemmanualassignMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

