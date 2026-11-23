package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitieslanguageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitieslanguageDud struct { 
    


    


    


    


    


    


    

}

// Capabilitieslanguage
type Capabilitieslanguage struct { 
    // IsAvailable
    IsAvailable bool `json:"isAvailable"`


    // UnavailableReason
    UnavailableReason Capabilitiesunavailablereason `json:"unavailableReason"`


    // Contexts
    Contexts Capabilitieslanguagecontextsmodel `json:"contexts"`


    // Fallback
    Fallback string `json:"fallback"`


    // Features
    Features Capabilitieslanguagefeatures `json:"features"`


    // Functions
    Functions map[string]Capabilitiesfunction `json:"functions"`


    // IsDefault
    IsDefault bool `json:"isDefault"`

}

// String returns a JSON representation of the model
func (o *Capabilitieslanguage) String() string {
    
    
    
    
    
     o.Functions = map[string]Capabilitiesfunction{"": {}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitieslanguage) MarshalJSON() ([]byte, error) {
    type Alias Capabilitieslanguage

    if CapabilitieslanguageMarshalled {
        return []byte("{}"), nil
    }
    CapabilitieslanguageMarshalled = true

    return json.Marshal(&struct {
        
        IsAvailable bool `json:"isAvailable"`
        
        UnavailableReason Capabilitiesunavailablereason `json:"unavailableReason"`
        
        Contexts Capabilitieslanguagecontextsmodel `json:"contexts"`
        
        Fallback string `json:"fallback"`
        
        Features Capabilitieslanguagefeatures `json:"features"`
        
        Functions map[string]Capabilitiesfunction `json:"functions"`
        
        IsDefault bool `json:"isDefault"`
        *Alias
    }{

        


        


        


        


        


        
        Functions: map[string]Capabilitiesfunction{"": {}},
        


        

        Alias: (*Alias)(u),
    })
}

