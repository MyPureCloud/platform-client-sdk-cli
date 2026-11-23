package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitieslanguagettsengineMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitieslanguagettsengineDud struct { 
    


    


    


    

}

// Capabilitieslanguagettsengine
type Capabilitieslanguagettsengine struct { 
    // FlowTypes
    FlowTypes map[string]Capabilitiesavailability `json:"flowTypes"`


    // Id
    Id string `json:"id"`


    // IsDefault
    IsDefault bool `json:"isDefault"`


    // Voices
    Voices map[string]Capabilitieslanguagettsenginevoice `json:"voices"`

}

// String returns a JSON representation of the model
func (o *Capabilitieslanguagettsengine) String() string {
     o.FlowTypes = map[string]Capabilitiesavailability{"": {}} 
    
    
     o.Voices = map[string]Capabilitieslanguagettsenginevoice{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitieslanguagettsengine) MarshalJSON() ([]byte, error) {
    type Alias Capabilitieslanguagettsengine

    if CapabilitieslanguagettsengineMarshalled {
        return []byte("{}"), nil
    }
    CapabilitieslanguagettsengineMarshalled = true

    return json.Marshal(&struct {
        
        FlowTypes map[string]Capabilitiesavailability `json:"flowTypes"`
        
        Id string `json:"id"`
        
        IsDefault bool `json:"isDefault"`
        
        Voices map[string]Capabilitieslanguagettsenginevoice `json:"voices"`
        *Alias
    }{

        
        FlowTypes: map[string]Capabilitiesavailability{"": {}},
        


        


        


        
        Voices: map[string]Capabilitieslanguagettsenginevoice{"": {}},
        

        Alias: (*Alias)(u),
    })
}

