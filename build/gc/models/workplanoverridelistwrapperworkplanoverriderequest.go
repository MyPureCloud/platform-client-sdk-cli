package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkplanoverridelistwrapperworkplanoverriderequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkplanoverridelistwrapperworkplanoverriderequestDud struct { 
    


    

}

// Workplanoverridelistwrapperworkplanoverriderequest
type Workplanoverridelistwrapperworkplanoverriderequest struct { 
    // Values
    Values []Workplanoverriderequest `json:"values"`


    // DeleteAll - if true, should delete all existing overrides for the agent and update the given overrides
    DeleteAll bool `json:"deleteAll"`

}

// String returns a JSON representation of the model
func (o *Workplanoverridelistwrapperworkplanoverriderequest) String() string {
     o.Values = []Workplanoverriderequest{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workplanoverridelistwrapperworkplanoverriderequest) MarshalJSON() ([]byte, error) {
    type Alias Workplanoverridelistwrapperworkplanoverriderequest

    if WorkplanoverridelistwrapperworkplanoverriderequestMarshalled {
        return []byte("{}"), nil
    }
    WorkplanoverridelistwrapperworkplanoverriderequestMarshalled = true

    return json.Marshal(&struct {
        
        Values []Workplanoverriderequest `json:"values"`
        
        DeleteAll bool `json:"deleteAll"`
        *Alias
    }{

        
        Values: []Workplanoverriderequest{{}},
        


        

        Alias: (*Alias)(u),
    })
}

