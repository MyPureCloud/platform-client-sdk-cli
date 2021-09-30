package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmergencycallflowMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmergencycallflowDud struct { 
    


    

}

// Emergencycallflow - An emergency flow associates a call flow to use in an emergency with the ivr(s) to route to it.
type Emergencycallflow struct { 
    // EmergencyFlow - The call flow to execute in an emergency.
    EmergencyFlow Domainentityref `json:"emergencyFlow"`


    // Ivrs - The IVR(s) to route to the call flow during an emergency.
    Ivrs []Domainentityref `json:"ivrs"`

}

// String returns a JSON representation of the model
func (o *Emergencycallflow) String() string {
    
    
    
    
    
    
     o.Ivrs = []Domainentityref{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Emergencycallflow) MarshalJSON() ([]byte, error) {
    type Alias Emergencycallflow

    if EmergencycallflowMarshalled {
        return []byte("{}"), nil
    }
    EmergencycallflowMarshalled = true

    return json.Marshal(&struct { 
        EmergencyFlow Domainentityref `json:"emergencyFlow"`
        
        Ivrs []Domainentityref `json:"ivrs"`
        
        *Alias
    }{
        

        

        

        
        Ivrs: []Domainentityref{{}},
        

        
        Alias: (*Alias)(u),
    })
}

