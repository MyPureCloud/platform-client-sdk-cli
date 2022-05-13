package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChatmediapolicyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChatmediapolicyDud struct { 
    


    

}

// Chatmediapolicy
type Chatmediapolicy struct { 
    // Actions - Actions applied when specified conditions are met
    Actions Policyactions `json:"actions"`


    // Conditions - Conditions for when actions should be applied
    Conditions Chatmediapolicyconditions `json:"conditions"`

}

// String returns a JSON representation of the model
func (o *Chatmediapolicy) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Chatmediapolicy) MarshalJSON() ([]byte, error) {
    type Alias Chatmediapolicy

    if ChatmediapolicyMarshalled {
        return []byte("{}"), nil
    }
    ChatmediapolicyMarshalled = true

    return json.Marshal(&struct {
        
        Actions Policyactions `json:"actions"`
        
        Conditions Chatmediapolicyconditions `json:"conditions"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

