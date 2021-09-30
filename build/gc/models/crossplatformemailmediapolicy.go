package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CrossplatformemailmediapolicyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CrossplatformemailmediapolicyDud struct { 
    


    

}

// Crossplatformemailmediapolicy
type Crossplatformemailmediapolicy struct { 
    // Actions - Actions applied when specified conditions are met
    Actions Crossplatformpolicyactions `json:"actions"`


    // Conditions - Conditions for when actions should be applied
    Conditions Emailmediapolicyconditions `json:"conditions"`

}

// String returns a JSON representation of the model
func (o *Crossplatformemailmediapolicy) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Crossplatformemailmediapolicy) MarshalJSON() ([]byte, error) {
    type Alias Crossplatformemailmediapolicy

    if CrossplatformemailmediapolicyMarshalled {
        return []byte("{}"), nil
    }
    CrossplatformemailmediapolicyMarshalled = true

    return json.Marshal(&struct { 
        Actions Crossplatformpolicyactions `json:"actions"`
        
        Conditions Emailmediapolicyconditions `json:"conditions"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

