package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesflowtypeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesflowtypeDud struct { 
    


    


    


    


    


    


    


    UserVariableLimits Capabilitiesflowtypeuservariablelimits `json:"userVariableLimits"`


    Limits Capabilitiesflowtypelimitinfo `json:"limits"`

}

// Capabilitiesflowtype
type Capabilitiesflowtype struct { 
    // IsAvailable
    IsAvailable bool `json:"isAvailable"`


    // UnavailableReason
    UnavailableReason Capabilitiesunavailablereason `json:"unavailableReason"`


    // Components
    Components map[string]Capabilitiesflowcomponent `json:"components"`


    // IsSecure
    IsSecure bool `json:"isSecure"`


    // SupportsLanguages
    SupportsLanguages Capabilitiesflowtypelanguagesupport `json:"supportsLanguages"`


    // UserCreatable
    UserCreatable Capabilitiesavailability `json:"userCreatable"`


    // MaxSize
    MaxSize int `json:"maxSize"`


    


    

}

// String returns a JSON representation of the model
func (o *Capabilitiesflowtype) String() string {
    
    
     o.Components = map[string]Capabilitiesflowcomponent{"": {}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiesflowtype) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiesflowtype

    if CapabilitiesflowtypeMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesflowtypeMarshalled = true

    return json.Marshal(&struct {
        
        IsAvailable bool `json:"isAvailable"`
        
        UnavailableReason Capabilitiesunavailablereason `json:"unavailableReason"`
        
        Components map[string]Capabilitiesflowcomponent `json:"components"`
        
        IsSecure bool `json:"isSecure"`
        
        SupportsLanguages Capabilitiesflowtypelanguagesupport `json:"supportsLanguages"`
        
        UserCreatable Capabilitiesavailability `json:"userCreatable"`
        
        MaxSize int `json:"maxSize"`
        *Alias
    }{

        


        


        
        Components: map[string]Capabilitiesflowcomponent{"": {}},
        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

