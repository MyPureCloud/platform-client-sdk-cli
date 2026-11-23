package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesflowcomponentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesflowcomponentDud struct { 
    


    


    

}

// Capabilitiesflowcomponent
type Capabilitiesflowcomponent struct { 
    // IsAvailable
    IsAvailable bool `json:"isAvailable"`


    // UnavailableReason
    UnavailableReason Capabilitiesunavailablereason `json:"unavailableReason"`


    // CanBeStartupObject
    CanBeStartupObject bool `json:"canBeStartupObject"`

}

// String returns a JSON representation of the model
func (o *Capabilitiesflowcomponent) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiesflowcomponent) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiesflowcomponent

    if CapabilitiesflowcomponentMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesflowcomponentMarshalled = true

    return json.Marshal(&struct {
        
        IsAvailable bool `json:"isAvailable"`
        
        UnavailableReason Capabilitiesunavailablereason `json:"unavailableReason"`
        
        CanBeStartupObject bool `json:"canBeStartupObject"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

