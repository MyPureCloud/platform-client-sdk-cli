package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitieslimitconditionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitieslimitconditionDud struct { 
    IsAvailable bool `json:"isAvailable"`


    Value Capabilitieslimitvalue `json:"value"`


    FlowTypes map[string]Capabilitieslimitcondition `json:"flowTypes"`

}

// Capabilitieslimitcondition
type Capabilitieslimitcondition struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Capabilitieslimitcondition) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitieslimitcondition) MarshalJSON() ([]byte, error) {
    type Alias Capabilitieslimitcondition

    if CapabilitieslimitconditionMarshalled {
        return []byte("{}"), nil
    }
    CapabilitieslimitconditionMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

