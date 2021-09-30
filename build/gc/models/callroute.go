package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CallrouteMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CallrouteDud struct { 
    

}

// Callroute
type Callroute struct { 
    // Targets - A list of CallTargets to be called when the CallRoute is executed
    Targets []Calltarget `json:"targets"`

}

// String returns a JSON representation of the model
func (o *Callroute) String() string {
    
    
     o.Targets = []Calltarget{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Callroute) MarshalJSON() ([]byte, error) {
    type Alias Callroute

    if CallrouteMarshalled {
        return []byte("{}"), nil
    }
    CallrouteMarshalled = true

    return json.Marshal(&struct { 
        Targets []Calltarget `json:"targets"`
        
        *Alias
    }{
        

        
        Targets: []Calltarget{{}},
        

        
        Alias: (*Alias)(u),
    })
}

