package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagemediapolicyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagemediapolicyDud struct { 
    


    

}

// Messagemediapolicy
type Messagemediapolicy struct { 
    // Actions - Actions applied when specified conditions are met
    Actions Policyactions `json:"actions"`


    // Conditions - Conditions for when actions should be applied
    Conditions Messagemediapolicyconditions `json:"conditions"`

}

// String returns a JSON representation of the model
func (o *Messagemediapolicy) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagemediapolicy) MarshalJSON() ([]byte, error) {
    type Alias Messagemediapolicy

    if MessagemediapolicyMarshalled {
        return []byte("{}"), nil
    }
    MessagemediapolicyMarshalled = true

    return json.Marshal(&struct {
        
        Actions Policyactions `json:"actions"`
        
        Conditions Messagemediapolicyconditions `json:"conditions"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

