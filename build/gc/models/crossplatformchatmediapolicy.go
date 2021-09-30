package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CrossplatformchatmediapolicyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CrossplatformchatmediapolicyDud struct { 
    


    

}

// Crossplatformchatmediapolicy
type Crossplatformchatmediapolicy struct { 
    // Actions - Actions applied when specified conditions are met
    Actions Crossplatformpolicyactions `json:"actions"`


    // Conditions - Conditions for when actions should be applied
    Conditions Chatmediapolicyconditions `json:"conditions"`

}

// String returns a JSON representation of the model
func (o *Crossplatformchatmediapolicy) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Crossplatformchatmediapolicy) MarshalJSON() ([]byte, error) {
    type Alias Crossplatformchatmediapolicy

    if CrossplatformchatmediapolicyMarshalled {
        return []byte("{}"), nil
    }
    CrossplatformchatmediapolicyMarshalled = true

    return json.Marshal(&struct { 
        Actions Crossplatformpolicyactions `json:"actions"`
        
        Conditions Chatmediapolicyconditions `json:"conditions"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

