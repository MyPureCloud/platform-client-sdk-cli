package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesdependencytypeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesdependencytypeDud struct { 
    


    

}

// Capabilitiesdependencytype
type Capabilitiesdependencytype struct { 
    // Availability
    Availability Capabilitiesavailability `json:"availability"`


    // IsVersioned
    IsVersioned bool `json:"isVersioned"`

}

// String returns a JSON representation of the model
func (o *Capabilitiesdependencytype) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiesdependencytype) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiesdependencytype

    if CapabilitiesdependencytypeMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesdependencytypeMarshalled = true

    return json.Marshal(&struct {
        
        Availability Capabilitiesavailability `json:"availability"`
        
        IsVersioned bool `json:"isVersioned"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

