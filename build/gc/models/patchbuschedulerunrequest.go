package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchbuschedulerunrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchbuschedulerunrequestDud struct { 
    

}

// Patchbuschedulerunrequest
type Patchbuschedulerunrequest struct { 
    // ReschedulingOptions - The rescheduling options to update
    ReschedulingOptions Patchbureschedulingoptionsrequest `json:"reschedulingOptions"`

}

// String returns a JSON representation of the model
func (o *Patchbuschedulerunrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Patchbuschedulerunrequest) MarshalJSON() ([]byte, error) {
    type Alias Patchbuschedulerunrequest

    if PatchbuschedulerunrequestMarshalled {
        return []byte("{}"), nil
    }
    PatchbuschedulerunrequestMarshalled = true

    return json.Marshal(&struct {
        
        ReschedulingOptions Patchbureschedulingoptionsrequest `json:"reschedulingOptions"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

