package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CallmediapolicyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CallmediapolicyDud struct { 
    


    

}

// Callmediapolicy
type Callmediapolicy struct { 
    // Actions - Actions applied when specified conditions are met
    Actions Policyactions `json:"actions"`


    // Conditions - Conditions for when actions should be applied
    Conditions Callmediapolicyconditions `json:"conditions"`

}

// String returns a JSON representation of the model
func (o *Callmediapolicy) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Callmediapolicy) MarshalJSON() ([]byte, error) {
    type Alias Callmediapolicy

    if CallmediapolicyMarshalled {
        return []byte("{}"), nil
    }
    CallmediapolicyMarshalled = true

    return json.Marshal(&struct {
        
        Actions Policyactions `json:"actions"`
        
        Conditions Callmediapolicyconditions `json:"conditions"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

