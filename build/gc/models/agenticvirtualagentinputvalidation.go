package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentinputvalidationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentinputvalidationDud struct { 
    

}

// Agenticvirtualagentinputvalidation
type Agenticvirtualagentinputvalidation struct { 
    // VarType - Validation type discriminator.
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentinputvalidation) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentinputvalidation) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentinputvalidation

    if AgenticvirtualagentinputvalidationMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentinputvalidationMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

