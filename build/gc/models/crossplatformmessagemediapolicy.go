package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CrossplatformmessagemediapolicyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CrossplatformmessagemediapolicyDud struct { 
    


    

}

// Crossplatformmessagemediapolicy
type Crossplatformmessagemediapolicy struct { 
    // Actions - Actions applied when specified conditions are met
    Actions Crossplatformpolicyactions `json:"actions"`


    // Conditions - Conditions for when actions should be applied
    Conditions Messagemediapolicyconditions `json:"conditions"`

}

// String returns a JSON representation of the model
func (o *Crossplatformmessagemediapolicy) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Crossplatformmessagemediapolicy) MarshalJSON() ([]byte, error) {
    type Alias Crossplatformmessagemediapolicy

    if CrossplatformmessagemediapolicyMarshalled {
        return []byte("{}"), nil
    }
    CrossplatformmessagemediapolicyMarshalled = true

    return json.Marshal(&struct {
        
        Actions Crossplatformpolicyactions `json:"actions"`
        
        Conditions Messagemediapolicyconditions `json:"conditions"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

