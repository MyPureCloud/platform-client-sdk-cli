package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmailmediapolicyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmailmediapolicyDud struct { 
    


    

}

// Emailmediapolicy
type Emailmediapolicy struct { 
    // Actions - Actions applied when specified conditions are met
    Actions Policyactions `json:"actions"`


    // Conditions - Conditions for when actions should be applied
    Conditions Emailmediapolicyconditions `json:"conditions"`

}

// String returns a JSON representation of the model
func (o *Emailmediapolicy) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Emailmediapolicy) MarshalJSON() ([]byte, error) {
    type Alias Emailmediapolicy

    if EmailmediapolicyMarshalled {
        return []byte("{}"), nil
    }
    EmailmediapolicyMarshalled = true

    return json.Marshal(&struct { 
        Actions Policyactions `json:"actions"`
        
        Conditions Emailmediapolicyconditions `json:"conditions"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

