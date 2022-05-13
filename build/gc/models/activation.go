package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivationDud struct { 
    


    

}

// Activation
type Activation struct { 
    // VarType - Type of activation.
    VarType string `json:"type"`


    // DelayInSeconds - Activation delay time amount.
    DelayInSeconds int `json:"delayInSeconds"`

}

// String returns a JSON representation of the model
func (o *Activation) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activation) MarshalJSON() ([]byte, error) {
    type Alias Activation

    if ActivationMarshalled {
        return []byte("{}"), nil
    }
    ActivationMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        DelayInSeconds int `json:"delayInSeconds"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

