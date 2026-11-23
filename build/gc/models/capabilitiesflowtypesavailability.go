package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesflowtypesavailabilityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesflowtypesavailabilityDud struct { 
    

}

// Capabilitiesflowtypesavailability
type Capabilitiesflowtypesavailability struct { 
    // FlowTypes
    FlowTypes map[string]Capabilitiesavailability `json:"flowTypes"`

}

// String returns a JSON representation of the model
func (o *Capabilitiesflowtypesavailability) String() string {
     o.FlowTypes = map[string]Capabilitiesavailability{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiesflowtypesavailability) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiesflowtypesavailability

    if CapabilitiesflowtypesavailabilityMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesflowtypesavailabilityMarshalled = true

    return json.Marshal(&struct {
        
        FlowTypes map[string]Capabilitiesavailability `json:"flowTypes"`
        *Alias
    }{

        
        FlowTypes: map[string]Capabilitiesavailability{"": {}},
        

        Alias: (*Alias)(u),
    })
}

