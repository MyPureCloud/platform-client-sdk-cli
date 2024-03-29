package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CrossplatformcallmediapolicyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CrossplatformcallmediapolicyDud struct { 
    


    

}

// Crossplatformcallmediapolicy
type Crossplatformcallmediapolicy struct { 
    // Actions - Actions applied when specified conditions are met
    Actions Crossplatformpolicyactions `json:"actions"`


    // Conditions - Conditions for when actions should be applied
    Conditions Callmediapolicyconditions `json:"conditions"`

}

// String returns a JSON representation of the model
func (o *Crossplatformcallmediapolicy) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Crossplatformcallmediapolicy) MarshalJSON() ([]byte, error) {
    type Alias Crossplatformcallmediapolicy

    if CrossplatformcallmediapolicyMarshalled {
        return []byte("{}"), nil
    }
    CrossplatformcallmediapolicyMarshalled = true

    return json.Marshal(&struct {
        
        Actions Crossplatformpolicyactions `json:"actions"`
        
        Conditions Callmediapolicyconditions `json:"conditions"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

