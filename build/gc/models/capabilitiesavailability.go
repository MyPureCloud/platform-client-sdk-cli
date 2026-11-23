package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesavailabilityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesavailabilityDud struct { 
    


    

}

// Capabilitiesavailability
type Capabilitiesavailability struct { 
    // IsAvailable
    IsAvailable bool `json:"isAvailable"`


    // UnavailableReason
    UnavailableReason Capabilitiesunavailablereason `json:"unavailableReason"`

}

// String returns a JSON representation of the model
func (o *Capabilitiesavailability) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiesavailability) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiesavailability

    if CapabilitiesavailabilityMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesavailabilityMarshalled = true

    return json.Marshal(&struct {
        
        IsAvailable bool `json:"isAvailable"`
        
        UnavailableReason Capabilitiesunavailablereason `json:"unavailableReason"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

